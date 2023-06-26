package engine

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/activity"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
)

// 处理CallActivity元素
func handleCallActivity(
	ctx context.Context,
	pi engine_types.ProcessInstance,
	callActivityElement *activity.TCallActivity,
) (bool, error) {
	// 找到或者创建激活的消息
	var childPI engine_types.ProcessInstance
	var cactiv engine_types.BaseElement
	var err error
	if cactiv, err = pi.GetElementManager().FindOneCallActivity(ctx, callActivityElement.ID); err != nil {
		cactiv, err = createChildInstance(ctx, pi, callActivityElement)
		if err != nil {
			return false, err
		}
	}
	// 继续子流程
	childPI, err = ContinuePIByKey(ctx, cactiv.(engine_types.CallActivity).GetChildPIKey())
	if err != nil {
		return false, err
	}
	switch childPI.GetState() {
	case sepc_pi_types.Completed:
		if err := pi.GetElementManager().DeleteCallActivity(ctx, cactiv.GetKey()); err != nil {
			return false, err
		}
		return true, nil
	default:
		return false, nil
	}
}

// 创建子流程
func createChildInstance(
	ctx context.Context,
	pi engine_types.ProcessInstance,
	callActivityElement *activity.TCallActivity,
) (engine_types.BaseElement, error) {
	if !callActivityElement.HasCalledElement() {
		return nil, fmt.Errorf("未存在子流程信息")
	}
	cactiv, err := pi.GetElementManager().CreateCallActivity(ctx, callActivityElement)
	if err != nil {
		return nil, err
	}
	variables := make(map[string]any)
	if callActivityElement.GetCalledElement().GetPropagateAllChildVariables() {
		variables = pi.GetVariables()
	}
	childPI, err := CreatePIByID(callActivityElement.GetCalledElement().GetProcessID(), variables)
	if err != nil {
		return nil, err
	}
	if err := childPI.SetParentProcessInstanceKey(ctx, pi.GetKey()); err != nil {
		return nil, err
	}
	if err := cactiv.(engine_types.CallActivity).SetChildPIKey(ctx, childPI.GetKey()); err != nil {
		return nil, err
	}
	return cactiv, nil
}
