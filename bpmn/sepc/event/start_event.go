package event

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// startEvent *TStartEvent github.com/averyyan/bpmn-engine/bpmn/sepc/types.BaseElement
type TStartEvent struct {
	ID                  string                   `xml:"id,attr" json:"id"`                                       // 元素ID
	Name                string                   `xml:"name,attr" json:"name"`                                   // 元素名称
	Documentation       string                   `xml:"documentation,attr" json:"documentation"`                 // 元素说明
	IncomingAssociation []string                 `xml:"incoming" json:"incoming"`                                // 元素入Flow元素IDs
	OutgoingAssociation []string                 `xml:"outgoing" json:"outgoing"`                                // 元素出Flow元素IDs
	Properties          []*extensions.TPropertie `xml:"extensionElements>properties>property" json:"properties"` // 外部数据
}

func (startEvent *TStartEvent) GetDocumentation() string {
	return startEvent.Documentation
}

func (startEvent *TStartEvent) GetProperties() []sepc_types.Propertie {
	return extensions.Properties2Interface(startEvent.Properties)
}

func (startEvent *TStartEvent) GetID() string {
	return startEvent.ID
}

func (startEvent *TStartEvent) GetName() string {
	return startEvent.Name
}

func (startEvent *TStartEvent) GetIncomingAssociation() []string {
	return startEvent.IncomingAssociation
}

func (startEvent *TStartEvent) GetOutgoingAssociation() []string {
	return startEvent.OutgoingAssociation
}

func (startEvent *TStartEvent) GetType() sepc_element_types.ElementType {
	return sepc_element_types.StartEvent
}
