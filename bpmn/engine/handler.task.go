package engine

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

// 处理ServiceTask元素 非阻塞
func handleServiceTask(
	ctx context.Context,
	pi engine_types.ProcessInstance,
	element sepc_types.ServiceTask,
) (bool, error) {
	if !element.HasTaskDefinition() {
		return false, fmt.Errorf("任务定义未存在")
	}
	taskType := element.GetTaskDefinition()
	handler, err := GetServiceTaskHandler(taskType.GetTypeName())
	if err != nil {
		return false, err
	}
	if err := handler(&activatedActivity{
		pi:          pi,
		baseElement: element.(sepc_types.BaseElement),
		completeHandler: func() {
		},
		failHandler: func(reason string) {
			panic(fmt.Errorf(reason))
		},
	}); err != nil {
		return false, err
	}
	return true, nil
}
