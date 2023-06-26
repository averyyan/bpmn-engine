package gateway

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 消息事件网关
// eventBasedGateway *TEventBasedGateway github.com/averyyan/bpmn-engine/bpmn/sepc/types.BaseElement
type TEventBasedGateway struct {
	ID                  string                   `xml:"id,attr" json:"id"`                                       // 元素ID
	Name                string                   `xml:"name,attr" json:"name"`                                   // 元素名称
	Documentation       string                   `xml:"documentation,attr" json:"documentation"`                 // 元素说明
	IncomingAssociation []string                 `xml:"incoming" json:"incoming"`                                // 元素入Flow元素IDs
	OutgoingAssociation []string                 `xml:"outgoing" json:"outgoing"`                                // 元素出Flow元素IDs
	Properties          []*extensions.TPropertie `xml:"extensionElements>properties>property" json:"properties"` // 外部数据
}

func (eventBasedGateway *TEventBasedGateway) GetDocumentation() string {
	return eventBasedGateway.Documentation
}

func (eventBasedGateway *TEventBasedGateway) GetProperties() []sepc_types.Propertie {
	return extensions.Properties2Interface(eventBasedGateway.Properties)
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
