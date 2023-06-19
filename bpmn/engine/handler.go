package engine

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/flow"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 元素处理函数 返回元素出的序列流
func handleElement(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
	baseElement sepc_types.BaseElement,
) ([]*flow.TSequenceFlow, error) {
	fmt.Println("正在运行", baseElement.GetName()) // debug
	definitions, err := pi.GetDefinitions()
	if err != nil {
		return nil, err
	}
	nextFlows := findSequenceFlows(definitions.Process.SequenceFlows, baseElement.GetOutgoingAssociation())
	switch baseElement.GetType() {
	case sepc_element_types.StartEvent:
	case sepc_element_types.EndEvent:
		if err := handleEndEvent(ctx, state, pi); err != nil {
			return nil, err
		}
		if len(nextFlows) > 0 {
			return nil, fmt.Errorf("结束事件后不应有继续事件,请检查流程详情")
		}
	case sepc_element_types.ServiceTask:
		if err := handleServiceTask(ctx, state, pi, baseElement.(sepc_types.ServiceTaskElement)); err != nil {
			return nil, err
		}
	case sepc_element_types.ExclusiveGateway:
		_nextFlows, err := handleExclusiveGateway(nextFlows, pi.GetVariables(), baseElement.(sepc_types.ExclusiveGateway))
		if err != nil {
			return nil, err
		}
		nextFlows = _nextFlows
	default:
		return nil, fmt.Errorf("未支持元素")
	}
	return nextFlows, nil
}
