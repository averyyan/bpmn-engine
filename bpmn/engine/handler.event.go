package engine

import (
	"context"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 处理EndEvent元素
func handleEndEvent(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
) error {
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
	// 判断流程是否完结
	if len(activs) > 0 {
		return nil
	}
	// 流程已经完结
	if err := state.ProcessInstanceManager().SetCompleted(ctx, pi); err != nil {
		return err
	}
	return nil
}
