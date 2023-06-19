package gateway

import (
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 并行网关
type TParallelGateway struct {
	ID                  string   `xml:"id,attr" bson:"id"`
	Name                string   `xml:"name,attr" bson:"name"`
	IncomingAssociation []string `xml:"incoming" bson:"incoming"`
	OutgoingAssociation []string `xml:"outgoing" bson:"outgoing"`
}

func (parallelGateway *TParallelGateway) GetID() string {
	return parallelGateway.ID
}

func (parallelGateway *TParallelGateway) GetName() string {
	return parallelGateway.Name
}

func (parallelGateway *TParallelGateway) GetIncomingAssociation() []string {
	return parallelGateway.IncomingAssociation
}

func (parallelGateway *TParallelGateway) GetOutgoingAssociation() []string {
	return parallelGateway.OutgoingAssociation
}

func (parallelGateway *TParallelGateway) GetType() sepc_element_types.ElementType {
	return sepc_element_types.ParallelGateway
}
