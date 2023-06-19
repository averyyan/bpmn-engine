package gateway

import sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"

// 独占网关
type TExclusiveGateway struct {
	ID                  string   `xml:"id,attr"`      // 元素ID
	Name                string   `xml:"name,attr"`    // 元素名称
	IncomingAssociation []string `xml:"incoming"`     // 元素入Flow元素IDs
	OutgoingAssociation []string `xml:"outgoing"`     // 元素出Flow元素IDs
	Default             string   `xml:"default,attr"` // 默认的
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
