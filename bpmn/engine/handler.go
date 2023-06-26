package engine

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/activity"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/event"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 元素处理函数 返回元素出的序列流
func handleElement(
	ctx context.Context,
	pi engine_types.ProcessInstance,
	baseElement sepc_types.BaseElement,
) (bool, error) {
	fmt.Println("正在运行", baseElement.GetName()) // debug
	switch baseElement.GetType() {
	case sepc_element_types.StartEvent:
		return true, nil
	case sepc_element_types.EndEvent:
		if err := handleEndEvent(ctx, pi); err != nil {
			return false, err
		}
		return false, nil
	case sepc_element_types.ExclusiveGateway:
		return true, nil
	case sepc_element_types.EventBasedGateway:
		return true, nil
	case sepc_element_types.ParallelGateway:
		return handleParallelGateway(ctx, pi, baseElement), nil
	case sepc_element_types.MessageIntermediateCatchEvent:
		return handleMessageIntermediateCatchEvent(ctx, pi, baseElement.(*event.TIntermediateCatchEvent))
	case sepc_element_types.ServiceTask:
		return handleServiceTask(ctx, pi, baseElement.(sepc_types.ServiceTask))
	case sepc_element_types.CallActivity:
		return handleCallActivity(ctx, pi, baseElement.(*activity.TCallActivity))
	default:
		return false, fmt.Errorf("元素【%s】未被支持", baseElement.GetType())
	}
}
