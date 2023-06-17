package event

import sepc_types "github.com/averyyan/bpmn-engine/engine/sepc/type"

type EndEvent interface {
	sepc_types.BaseElement
}

type TEndEvent struct {
	ID                  string   `xml:"id,attr" bson:"id"`
	Name                string   `xml:"name,attr" bson:"name"`
	IncomingAssociation []string `xml:"incoming" bson:"incoming"`
	OutgoingAssociation []string `xml:"outgoing" bson:"outgoing"`
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
