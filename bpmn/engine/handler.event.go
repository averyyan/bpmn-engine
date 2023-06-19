package engine

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 处理EndEvent元素
func handleEndEvent(ctx context.Context, state engine_types.Engine, pi engine_types.ProcessInstance) error {
	// 检查未完成的任务
	activs, err := state.ActivityManager().FindByStates(
		ctx, pi,
		[]sepc_element_types.LifecycleState{
			sepc_element_types.Ready,
			sepc_element_types.Active,
		},
	)
	if err != nil {
		return err
	}
	if len(activs) > 0 {
		if err := state.ProcessInstanceManager().SetFailed(ctx, pi); err != nil {
			return err
		}
		return fmt.Errorf("存在未完成任务,请检查流程逻辑")
	}
	if err := state.ProcessInstanceManager().SetCompleted(ctx, pi); err != nil {
		return err
	}
	return nil
}
