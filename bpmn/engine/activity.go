package engine

import (
	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

// activ *activatedActivity github.com/averyyan/bpmn-engine/bpmn/engine/types.ActivatedActivity
type activatedActivity struct {
	pi              engine_types.ProcessInstance
	baseElement     sepc_types.BaseElement
	completeHandler func()
	failHandler     func(reason string)
}

// 获取流程实例
func (activ *activatedActivity) GetProcessInstance() engine_types.ProcessInstance {
	return activ.pi
}

func (activ *activatedActivity) GetElement() sepc_types.BaseElement {
	return activ.baseElement
}

func (activ *activatedActivity) Fail(reason string) {
	activ.failHandler(reason)
}

func (activ *activatedActivity) Complete() {
	activ.completeHandler()
}
