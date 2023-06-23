package gateway

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 并行网关
// parallelGateway *TParallelGateway github.com/averyyan/bpmn-engine/bpmn/sepc/types.BaseElement
type TParallelGateway struct {
	ID                  string                   `xml:"id,attr"`                               // 元素ID
	Name                string                   `xml:"name,attr"`                             // 元素名称
	Documentation       string                   `xml:"documentation,attr"`                    // 元素说明
	IncomingAssociation []string                 `xml:"incoming"`                              // 元素入Flow元素IDs
	OutgoingAssociation []string                 `xml:"outgoing"`                              // 元素出Flow元素IDs
	Properties          []*extensions.TPropertie `xml:"extensionElements>properties>property"` // 扩展数据
}

func (parallelGateway *TParallelGateway) GetDocumentation() string {
	return parallelGateway.Documentation
}

func (parallelGateway *TParallelGateway) GetProperties() []sepc_types.Propertie {
	return extensions.Properties2Interface(parallelGateway.Properties)
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
