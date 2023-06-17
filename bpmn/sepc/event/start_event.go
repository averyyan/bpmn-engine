package event

import sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"

type TStartEvent struct {
	ID                  string   `xml:"id,attr"`   // 元素ID
	Name                string   `xml:"name,attr"` // 元素名称
	IncomingAssociation []string `xml:"incoming"`  // 元素入Flow元素IDs
	OutgoingAssociation []string `xml:"outgoing"`  // 元素出Flow元素IDs
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
