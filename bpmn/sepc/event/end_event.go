package event

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// endEvent *TEndEvent github.com/averyyan/bpmn-engine/bpmn/sepc/types.BaseElement
type TEndEvent struct {
	ID                  string                   `xml:"id,attr"`                               // 元素ID
	Name                string                   `xml:"name,attr"`                             // 元素名称
	Documentation       string                   `xml:"documentation,attr"`                    // 元素说明
	IncomingAssociation []string                 `xml:"incoming"`                              // 元素入Flow元素IDs
	OutgoingAssociation []string                 `xml:"outgoing"`                              // 元素出Flow元素IDs
	Properties          []*extensions.TPropertie `xml:"extensionElements>properties>property"` // 扩展数据
}

func (endEvent *TEndEvent) GetDocumentation() string {
	return endEvent.Documentation
}

func (endEvent *TEndEvent) GetProperties() []sepc_types.Propertie {
	return extensions.Properties2Interface(endEvent.Properties)
}

func (endEvent *TEndEvent) GetID() string {
	return endEvent.ID
}

func (endEvent *TEndEvent) GetName() string {
	return endEvent.Name
}

func (endEvent *TEndEvent) GetIncomingAssociation() []string {
	return endEvent.IncomingAssociation
}

func (endEvent *TEndEvent) GetOutgoingAssociation() []string {
	return endEvent.OutgoingAssociation
}

func (endEvent *TEndEvent) GetType() sepc_element_types.ElementType {
	return sepc_element_types.EndEvent
}
