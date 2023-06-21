package memory_engine

import (
	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	memory_activity_manager "github.com/averyyan/bpmn-engine/memory/activity/manager"
	memory_ice_manager "github.com/averyyan/bpmn-engine/memory/ice/manager"
	memory_msgsub_manager "github.com/averyyan/bpmn-engine/memory/msgsub/manager"
	memory_process_instance_manager "github.com/averyyan/bpmn-engine/memory/process_instance/manager"
	memory_task_handler_manager "github.com/averyyan/bpmn-engine/memory/task_handler/manager"
)

func New() *Engine {
	return &Engine{
		processInstanceManager: memory_process_instance_manager.New(),
		activityManager:        memory_activity_manager.New(),
		taskHandlerManager:     memory_task_handler_manager.New(),
		iceManager:             memory_ice_manager.New(),
		msgsubManager:          memory_msgsub_manager.New(),
	}
}

// engine *Engine github.com/averyyan/bpmn-engine/bpmn/engine/types.Engine
type Engine struct {
	processInstanceManager engine_types.ProcessInstanceManager
	activityManager        engine_types.ActivityManager
	taskHandlerManager     engine_types.TaskHandlerManager
	iceManager             engine_types.ICEManager
	msgsubManager          engine_types.MessageSubscriptionManager
}

func (engine *Engine) ICEManager() engine_types.ICEManager {
	return engine.iceManager
}

func (engine *Engine) MessageSubscriptionManager() engine_types.MessageSubscriptionManager {
	return engine.msgsubManager
}

func (engine *Engine) ActivityManager() engine_types.ActivityManager {
	return engine.activityManager
}

func (engine *Engine) TaskHandlerManager() engine_types.TaskHandlerManager {
	return engine.taskHandlerManager
}

func (engine *Engine) ProcessInstanceManager() engine_types.ProcessInstanceManager {
	return engine.processInstanceManager
}
