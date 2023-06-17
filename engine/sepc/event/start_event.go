package event

import sepc_types "github.com/averyyan/bpmn-engine/engine/sepc/type"

type StartEvent interface {
	sepc_types.BaseElement
}

type TStartEvent struct {
	ID                  string   `xml:"id,attr"`
	Name                string   `xml:"name,attr"`
	IsInterrupting      bool     `xml:"isInterrupting,attr"`
	ParallelMultiple    bool     `xml:"parallelMultiple,attr"`
	IncomingAssociation []string `xml:"incoming"`
	OutgoingAssociation []string `xml:"outgoing"`
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
