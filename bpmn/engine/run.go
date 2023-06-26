package engine

import (
	"context"
	"fmt"
	"runtime/debug"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
)

func handlePanic(ctx context.Context, pi engine_types.ProcessInstance) {
	if r := recover(); r != nil {
		fmt.Println("流程实例运行错误:" + r.(error).Error() + string(debug.Stack()))
		if err := pi.SetFailed(ctx); err != nil {
			fmt.Println("流程实例无法操作")
		}
	}
}

// 流程运行核心
func run(ctx context.Context, pi engine_types.ProcessInstance) error {
	defer handlePanic(ctx, pi)
	if pi == nil {
		panic(fmt.Errorf("未存在流程实例"))
	}
	definitions, err := pi.GetDefinitions()
	if err != nil {
		return err
	}
	elementMap := makeElementMap(definitions)
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
		if err := pi.SetActive(ctx); err != nil {
			return err
		}
	case sepc_pi_types.Active:
		for _, element := range pi.GetElementManager().FindActiveElements(ctx) {
			queue = append(queue, queueElement{
				inboundFlowId: "",
				baseElement:   element,
			})
		}
	case sepc_pi_types.Completed:
	case sepc_pi_types.Failed:
	default:
		return fmt.Errorf("未支持【%s】流程实例状态", pi.GetState())
	}
	for len(queue) > 0 {
		element := queue[0].baseElement
		inboundFlowId := queue[0].inboundFlowId
		queue = queue[1:]
		continueNextElement, err := handleElement(ctx, pi, element)
		if err != nil {
			// 如果元素处理错误则将流程状态设置为失败
			if err := pi.SetFailed(ctx); err != nil {
				return err
			}
			return err
		}
		if continueNextElement {
			if inboundFlowId != "" {
				if err := pi.RemoveScheduledFlow(ctx, inboundFlowId); err != nil {
					return err
				}
			}
			nextFlows := findSequenceFlows(definitions.Process.SequenceFlows, element.GetOutgoingAssociation())
			if element.GetType() == sepc_element_types.ExclusiveGateway {
				_nextFlows, err := exclusivelyFilterByConditionExpression(nextFlows, pi.GetVariables(), element.(sepc_types.ExclusiveGateway))
				if err != nil {
					return err
				}
				nextFlows = _nextFlows
			}
			for _, nextFlow := range nextFlows {
				if err := pi.AppendScheduledFlow(ctx, nextFlow.ID); err != nil {
					return err
				}
				if baseElement, ok := elementMap[nextFlow.TargetRef]; ok {
					queue = append(queue, queueElement{
						inboundFlowId: nextFlow.ID,
						baseElement:   baseElement,
					})
				} else {
					return fmt.Errorf("元素【%s】未找到", nextFlow.ID)
				}
			}
		}
	}
	return nil
}
