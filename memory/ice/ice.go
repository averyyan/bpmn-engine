package ice

import (
	"time"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	"github.com/segmentio/ksuid"
)

func New(pi engine_types.ProcessInstance, baseElement sepc_types.BaseElement) *IntermediateCatchEvent {
	return &IntermediateCatchEvent{
		elementID:          baseElement.GetID(),
		elementName:        baseElement.GetName(),
		key:                ksuid.New().String(),
		processInstanceKey: pi.GetKey(),
		createTime:         time.Now().Local(),
		state:              sepc_element_types.Active,
	}
}

// ice *IntermediateCatchEvent github.com/averyyan/bpmn-engine/bpmn/engine/types.BaseElement
type IntermediateCatchEvent struct {
	elementID          string
	elementName        string
	key                string
	processInstanceKey string
	createTime         time.Time
	state              sepc_element_types.LifecycleState
}

func (ice *IntermediateCatchEvent) GetElementID() string {
	return ice.elementID
}

func (ice *IntermediateCatchEvent) GetElementName() string {
	return ice.elementName
}

func (ice *IntermediateCatchEvent) GetKey() string {
	return ice.key
}

func (ice *IntermediateCatchEvent) GetProcessInstanceKey() string {
	return ice.processInstanceKey
}

func (ice *IntermediateCatchEvent) GetCreateTime() time.Time {
	return ice.createTime
}

func (ice *IntermediateCatchEvent) GetState() sepc_element_types.LifecycleState {
	return ice.state
}

func (ice *IntermediateCatchEvent) SetState(state sepc_element_types.LifecycleState) {
	ice.state = state
}
