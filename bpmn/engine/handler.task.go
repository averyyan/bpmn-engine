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
	element sepc_types.ServiceTask,
) (bool, error) {
	// 创建一个新的活动
	activ, err := state.ActivityManager().(engine_types.BaseManager).Create(ctx, pi, element.(sepc_types.BaseElement))
	if err != nil {
		return false, err
	}
	if !element.HasTaskDefinition() {
		return false, fmt.Errorf("任务定义未存在")
	}
	taskType := element.GetTaskDefinition()
	handler, err := state.TaskHandlerManager().GetServiceTaskHandler(taskType.GetTypeName())
	if err != nil {
		return false, err
	}
	if err := handler(&activatedActivity{
		pi:          pi,
		activ:       activ,
		baseElement: element.(sepc_types.BaseElement),
		completeHandler: func() error {
			if err := state.ActivityManager().(engine_types.BaseManager).SetCompleted(ctx, activ); err != nil {
				return err
			}
			return nil
		},
		failHandler: func(reason string) error {
			if err := state.ActivityManager().(engine_types.BaseManager).SetFailed(ctx, activ, reason); err != nil {
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
