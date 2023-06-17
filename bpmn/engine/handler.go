package engine

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

func handleElement(ctx context.Context, state engine_types.Engine, pi engine_types.ProcessInstance, baseElement sepc_types.BaseElement) (bool, error) {
	fmt.Println("正在运行", baseElement.GetName()) // debug
	switch baseElement.GetType() {
	case sepc_element_types.StartEvent:
		return true, nil
	case sepc_element_types.EndEvent:
		return handleEndEvent(ctx, state, pi)
	}
	return true, nil
}

// 处理EndEvent元素
func handleEndEvent(ctx context.Context, state engine_types.Engine, pi engine_types.ProcessInstance) (bool, error) {
	if err := state.ProcessInstanceManager().SetCompleted(ctx, pi); err != nil {
		return false, err
	}
	return false, nil
}
