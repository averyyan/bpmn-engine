package engine

import (
	"context"
	"fmt"
	"os"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
)

// 根据BPMN文件地址创建流程实例
func CreateInstanceByFileAndRun(ctx context.Context, state engine_types.Engine, filepath string, variables map[string]any) (engine_types.ProcessInstance, error) {
	raw, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	pi, err := state.ProcessInstanceManager().Create(ctx, raw, variables)
	if err != nil {
		return nil, err
	}
	return pi, run(ctx, state, pi)
}

// 流程运行核心
func run(ctx context.Context, state engine_types.Engine, pi engine_types.ProcessInstance) error {
	definitions, err := pi.GetDefinitions()
	if err != nil {
		return err
	}
	// 内部元素队列
	type queueElement struct {
		inboundFlowId string
		baseElement   sepc_types.BaseElement
	}
	queue := make([]queueElement, 0)
	switch pi.GetState() {
	case sepc_pi_types.Ready:
		// 将开始事件排入待处理元素队列
		for _, event := range definitions.Process.StartEvents {
			queue = append(queue, queueElement{
				inboundFlowId: "",
				baseElement:   event,
			})
		}
		// 设置流程实例为激活状态
		if err := state.ProcessInstanceManager().SetActive(ctx, pi); err != nil {
			return err
		}
	default:
		return fmt.Errorf("未存在此流程实例状态")
	}
	// 循环元素队列直到完成 TODO:协程并行
	for len(queue) > 0 {
		baseElement := queue[0].baseElement
		inboundFlowId := queue[0].inboundFlowId
		// 核心函数 处理元素并且判断是否存在下一个元素
		nextFlows, err := handleElement(ctx, state, pi, baseElement)
		if err != nil {
			return err
		}
		if len(nextFlows) > 0 {
			if inboundFlowId != "" {
				if err := state.ProcessInstanceManager().RemoveScheduledFlows(ctx, pi, inboundFlowId); err != nil {
					return err
				}
			}
			for _, flow := range nextFlows {
				if err := state.ProcessInstanceManager().AppendScheduledFlows(ctx, pi, flow.ID); err != nil {
					return err
				}
				baseElements, err := pi.FindBaseElementsById(flow.TargetRef)
				if err != nil {
					return err
				}
				queue = append(queue, queueElement{
					inboundFlowId: flow.ID,
					baseElement:   baseElements,
				})
			}
		}
		// 删除已处理元素
		queue = queue[1:]
	}
	return nil
}
