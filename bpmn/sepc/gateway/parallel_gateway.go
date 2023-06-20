package gateway

import (
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 并行网关
type TParallelGateway struct {
	ID                  string   `xml:"id,attr"`   // 元素ID
	Name                string   `xml:"name,attr"` // 元素名称
	IncomingAssociation []string `xml:"incoming"`  // 元素入Flow元素IDs
	OutgoingAssociation []string `xml:"outgoing"`  // 元素出Flow元素IDs
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
