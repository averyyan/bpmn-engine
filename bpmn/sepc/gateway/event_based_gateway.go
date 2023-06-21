package gateway

import sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"

// 消息事件网关
type TEventBasedGateway struct {
	ID                  string   `xml:"id,attr"`   // 元素ID
	Name                string   `xml:"name,attr"` // 元素名称
	IncomingAssociation []string `xml:"incoming"`  // 元素入Flow元素IDs
	OutgoingAssociation []string `xml:"outgoing"`  // 元素出Flow元素IDs
}

func (eventBasedGateway *TEventBasedGateway) GetID() string {
	return eventBasedGateway.ID
}

func (eventBasedGateway *TEventBasedGateway) GetName() string {
	return eventBasedGateway.Name
}

func (eventBasedGateway *TEventBasedGateway) GetIncomingAssociation() []string {
	return eventBasedGateway.IncomingAssociation
}

func (eventBasedGateway *TEventBasedGateway) GetOutgoingAssociation() []string {
	return eventBasedGateway.OutgoingAssociation
}

func (eventBasedGateway *TEventBasedGateway) GetType() sepc_element_types.ElementType {
	return sepc_element_types.EventBasedGateway
}
