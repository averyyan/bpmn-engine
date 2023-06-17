package memory_engine

import (
	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	memory_process_instance_manager "github.com/averyyan/bpmn-engine/memory/process_instance/manager"
)

func New() *Engine {
	return &Engine{
		processInstanceManager: memory_process_instance_manager.New(),
	}
}

// engine *Engine github.com/averyyan/bpmn-engine/bpmn/engine/types.Engine
type Engine struct {
	processInstanceManager engine_types.ProcessInstanceManager
}

func (engine *Engine) ProcessInstanceManager() engine_types.ProcessInstanceManager {
	return engine.processInstanceManager
}
