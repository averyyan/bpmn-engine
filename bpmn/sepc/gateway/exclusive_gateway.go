package gateway

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 独占网关
// exclusiveGateway *TExclusiveGateway github.com/averyyan/bpmn-engine/bpmn/sepc/types.BaseElement
// exclusiveGateway *TExclusiveGateway github.com/averyyan/bpmn-engine/bpmn/sepc/types.ExclusiveGateway
type TExclusiveGateway struct {
	ID                  string                   `xml:"id,attr" json:"id"`                                       // 元素ID
	Name                string                   `xml:"name,attr" json:"name"`                                   // 元素名称
	Documentation       string                   `xml:"documentation,attr" json:"documentation"`                 // 元素说明
	IncomingAssociation []string                 `xml:"incoming" json:"incoming"`                                // 元素入Flow元素IDs
	OutgoingAssociation []string                 `xml:"outgoing" json:"outgoing"`                                // 元素出Flow元素IDs
	Properties          []*extensions.TPropertie `xml:"extensionElements>properties>property" json:"properties"` // 外部数据
	Default             string                   `xml:"default,attr" json:"default"`                             // 默认的线路的ID
}

func (exclusiveGateway *TExclusiveGateway) GetDocumentation() string {
	return exclusiveGateway.Documentation
}

func (exclusiveGateway *TExclusiveGateway) GetProperties() []sepc_types.Propertie {
	return extensions.Properties2Interface(exclusiveGateway.Properties)
}

func (exclusiveGateway *TExclusiveGateway) GetID() string {
	return exclusiveGateway.ID
}

func (exclusiveGateway *TExclusiveGateway) GetName() string {
	return exclusiveGateway.Name
}

func (exclusiveGateway *TExclusiveGateway) GetIncomingAssociation() []string {
	return exclusiveGateway.IncomingAssociation
}

func (exclusiveGateway *TExclusiveGateway) GetOutgoingAssociation() []string {
	return exclusiveGateway.OutgoingAssociation
}

func (exclusiveGateway *TExclusiveGateway) GetType() sepc_element_types.ElementType {
	return sepc_element_types.ExclusiveGateway
}

func (exclusiveGateway *TExclusiveGateway) GetDefault() string {
	return exclusiveGateway.Default
}
