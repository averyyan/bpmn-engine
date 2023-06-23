package engine

import (
	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

// activ *activatedActivity github.com/averyyan/bpmn-engine/bpmn/engine/types.ActivatedActivity
type activatedActivity struct {
	pi              engine_types.ProcessInstance
	activ           engine_types.Activity
	baseElement     sepc_types.BaseElement
	completeHandler func() error
	failHandler     func(reason string) error
}

// 获取流程实例
func (activ *activatedActivity) GetProcessInstance() engine_types.ProcessInstance {
	return activ.pi
}

func (activ *activatedActivity) GetActivity() engine_types.Activity {
	return activ.activ
}

func (activ *activatedActivity) GetElement() sepc_types.ServiceTask {
	return activ.baseElement.(sepc_types.ServiceTask)
}

func (activ *activatedActivity) Fail(reason string) error {
	return activ.failHandler(reason)
}

func (activ *activatedActivity) Complete() error {
	return activ.completeHandler()
}
