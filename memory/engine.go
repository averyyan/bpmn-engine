package memory_engine

import (
	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	memory_process_manager "github.com/averyyan/bpmn-engine/memory/process/manager"
	memory_process_instance_manager "github.com/averyyan/bpmn-engine/memory/process_instance/manager"
	memory_task_handler_manager "github.com/averyyan/bpmn-engine/memory/task_handler/manager"
)

func New() *Engine {
	return &Engine{
		EProcessManager:         memory_process_manager.New(),
		EProcessInstanceManager: memory_process_instance_manager.New(),
		taskHandlerManager:      memory_task_handler_manager.New(),
	}
}

// engine *Engine github.com/averyyan/bpmn-engine/bpmn/engine/types.Engine
type Engine struct {
	EProcessManager         *memory_process_manager.ProcessManager                  `json:"process_manager"`
	EProcessInstanceManager *memory_process_instance_manager.ProcessInstanceManager `json:"process_instance_manager"`
	taskHandlerManager      engine_types.TaskHandlerManager
}

func (engine *Engine) ProcessManager() engine_types.ProcessManager {
	return engine.EProcessManager
}

func (engine *Engine) ProcessInstanceManager() engine_types.ProcessInstanceManager {
	return engine.EProcessInstanceManager
}

// 服务任务管理
func (engine *Engine) TaskHandlerManager() engine_types.TaskHandlerManager {
	return engine.taskHandlerManager
}
