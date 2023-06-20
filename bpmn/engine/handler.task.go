package engine

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

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
		pi:          pi,
		activ:       activ,
		baseElement: element,
		completeHandler: func() error {
			if err := state.ActivityManager().SetCompleted(ctx, activ); err != nil {
				return err
			}
			return nil
		},
		failHandler: func(reason string) error {
			if err := state.ActivityManager().SetFailed(ctx, activ, reason); err != nil {
				return err
			}
			if err := state.ProcessInstanceManager().SetFailed(ctx, pi); err != nil {
				return err
			}
			return fmt.Errorf(reason)
		},
	}); err != nil {
		return false, err
	}
	return true, nil
}
