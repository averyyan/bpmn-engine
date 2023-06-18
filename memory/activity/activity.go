package activity

import (
	"time"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	"github.com/segmentio/ksuid"
)

func New(pi engine_types.ProcessInstance, baseElement sepc_types.BaseElement) *Activity {
	return &Activity{
		elementID:   baseElement.GetID(),
		elementName: baseElement.GetName(),
		key:         ksuid.New().String(),
		createTime:  time.Now().Local(),
		state:       sepc_element_types.Ready,
	}
}

// activ *Activity github.com/averyyan/bpmn-engine/bpmn/engine/types.BaseElement
type Activity struct {
	elementID          string
	elementName        string
	key                string
	processInstanceKey string
	createTime         time.Time
	state              sepc_element_types.LifecycleState
}

func (activ *Activity) GetElementID() string {
	return activ.elementID
}

func (activ *Activity) GetElementName() string {
	return activ.elementName
}

func (activ *Activity) GetKey() string {
	return activ.key
}

func (activ *Activity) GetProcessInstanceKey() string {
	return activ.processInstanceKey
}

func (activ *Activity) GetCreateTime() time.Time {
	return activ.createTime
}

func (activ *Activity) GetState() sepc_element_types.LifecycleState {
	return activ.state
}

func (activ *Activity) SetState(state sepc_element_types.LifecycleState) {
	activ.state = state
}
