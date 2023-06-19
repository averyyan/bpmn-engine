package gateway

import sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"

// 消息事件网关
type TEventBasedGateway struct {
	ID                  string   `xml:"id,attr" bson:"id"`
	Name                string   `xml:"name,attr" bson:"name"`
	IncomingAssociation []string `xml:"incoming" bson:"incoming"`
	OutgoingAssociation []string `xml:"outgoing" bson:"outgoing"`
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
