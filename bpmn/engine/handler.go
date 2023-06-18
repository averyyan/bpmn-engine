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
	case sepc_element_types.ServiceTask:
		return handleServiceTask(ctx, state, pi, baseElement.(sepc_types.ServiceTaskElement))
	}
	return true, nil
}

// 处理ServiceTask元素
func handleServiceTask(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
	element sepc_types.ServiceTaskElement,
) (bool, error) {
	// 创建一个新的活动
	activ, err := state.ActivityManager().Create(ctx, pi, element)
	if err != nil {
		return false, err
	}
	taskType := element.GetTaskDefinitionType()
	handler, err := state.TaskHandlerManager().GetServiceTaskHandler(taskType)
	if err != nil {
		return false, err
	}
	if err := handler(&activatedActivity{
		activ:       activ,
		baseElement: element,
		completeHandler: func() error {
			return state.ActivityManager().SetCompleted(ctx, activ)
		},
		failHandler: func(reason string) error {
			return state.ActivityManager().SetFailed(ctx, activ, reason)
		},
	}); err != nil {
		return false, err
	}
	return true, nil
}

// 处理EndEvent元素
func handleEndEvent(ctx context.Context, state engine_types.Engine, pi engine_types.ProcessInstance) (bool, error) {
	// 检查未完成的任务
	activs, err := state.ActivityManager().FindByStates(
		ctx, pi,
		[]sepc_element_types.LifecycleState{
			sepc_element_types.Ready,
			sepc_element_types.Active,
		},
	)
	if err != nil {
		return false, err
	}
	if len(activs) > 0 {
		if err := state.ProcessInstanceManager().SetFailed(ctx, pi); err != nil {
			return false, err
		}
		return false, fmt.Errorf("存在未完成任务,请检查流程逻辑")
	}
	if err := state.ProcessInstanceManager().SetCompleted(ctx, pi); err != nil {
		return false, err
	}
	return false, nil
}
