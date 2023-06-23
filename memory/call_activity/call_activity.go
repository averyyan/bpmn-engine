package call_activity

import (
	"time"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	"github.com/segmentio/ksuid"
)

func New(pi engine_types.ProcessInstance, baseElement sepc_types.BaseElement) *CallActivity {
	return &CallActivity{
		elementID:          baseElement.GetID(),
		elementName:        baseElement.GetName(),
		key:                ksuid.New().String(),
		processInstanceKey: pi.GetKey(),
		createTime:         time.Now().Local(),
		state:              sepc_element_types.Ready,
	}
}

// activ *CallActivity github.com/averyyan/bpmn-engine/bpmn/engine/types.BaseElement
// activ *CallActivity github.com/averyyan/bpmn-engine/bpmn/engine/types.CallActivity
type CallActivity struct {
	elementID               string
	elementName             string
	key                     string
	processInstanceKey      string
	createTime              time.Time
	state                   sepc_element_types.LifecycleState
	childProcessInstanceKey string
}

func (activ *CallActivity) SetChildProcessInstanceKey(key string) {
	activ.childProcessInstanceKey = key
}

func (activ *CallActivity) GetChildProcessInstanceKey() string {
	return activ.childProcessInstanceKey
}

func (activ *CallActivity) GetElementID() string {
	return activ.elementID
}

func (activ *CallActivity) GetElementName() string {
	return activ.elementName
}

func (activ *CallActivity) GetKey() string {
	return activ.key
}

func (activ *CallActivity) GetProcessInstanceKey() string {
	return activ.processInstanceKey
}

func (activ *CallActivity) GetCreateTime() time.Time {
	return activ.createTime
}

func (activ *CallActivity) SetState(state sepc_element_types.LifecycleState) {
	activ.state = state
}

func (activ *CallActivity) GetState() sepc_element_types.LifecycleState {
	return activ.state
}
