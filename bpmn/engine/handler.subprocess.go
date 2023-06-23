package engine

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
)

// 处理CallActivity元素
func handleCallActivity(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
	callActivityElement sepc_types.CallActivity,
) (bool, error) {
	if !callActivityElement.HasCalledElement() {
		return false, fmt.Errorf("未存在子流程信息")
	}
	var childPI engine_types.ProcessInstance
	var cactiv engine_types.BaseElement
	var err error
	// 找到活动的CallActivity
	cactiv, err = state.CallActivityManager().(engine_types.BaseManager).FindOneByStateAndID(
		ctx, pi, sepc_element_types.Active,
		callActivityElement.(sepc_types.BaseElement).GetID(),
	)
	if err != nil {
		cactiv, err = createChildInstance(ctx, state, pi, callActivityElement)
		if err != nil {
			return false, err
		}
	}
	// 继续子流程
	childPIkey := cactiv.(engine_types.CallActivity).GetChildProcessInstanceKey()
	childPI, err = ContinueProcessInstance(ctx, state, childPIkey)
	if err != nil {
		return false, err
	}
	switch childPI.GetState() {
	case sepc_pi_types.Completed:
		if err := state.CallActivityManager().(engine_types.BaseManager).SetCompleted(
			ctx, cactiv,
		); err != nil {
			return false, nil
		}
		return true, nil
	default:
		return false, nil
	}
}
