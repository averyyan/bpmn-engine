package memory_task_handler_manager

import (
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
)

func New() *TaskHandlerManager {
	return &TaskHandlerManager{
		handlers: make(map[string]engine_types.TaskHandler),
	}
}

// manager *TaskHandlerManager github.com/averyyan/bpmn-engine/bpmn/engine/types.TaskHandlerManager
type TaskHandlerManager struct {
	handlers map[string]engine_types.TaskHandler
}

// 注册 Servcie Task 元素处理函数
func (manager *TaskHandlerManager) RegisterServiceTaskHandler(taskType string, fun engine_types.TaskHandler) {
	manager.handlers[taskType] = fun
}

// 获取 Servcie Task 元素处理函数
func (manager *TaskHandlerManager) GetServiceTaskHandler(taskType string) (engine_types.TaskHandler, error) {
	handler, ok := manager.handlers[taskType]
	if !ok {
		return nil, fmt.Errorf("未存在元素【%s】绑定函数", taskType)
	}
	return handler, nil
}
