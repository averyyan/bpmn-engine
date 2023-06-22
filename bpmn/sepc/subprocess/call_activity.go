package subprocess

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

type TCallActivity struct {
	ID                  string                    `xml:"id,attr"`   // 元素ID
	Name                string                    `xml:"name,attr"` // 元素名称
	IncomingAssociation []string                  `xml:"incoming"`  // 元素入Flow元素IDs
	OutgoingAssociation []string                  `xml:"outgoing"`  // 元素出Flow元素IDs
	CalledElement       extensions.TCalledElement `xml:"extensionElements>calledElement" bson:"called_element"`
	Properties          []extensions.TPropertie   `xml:"extensionElements>properties>property" bson:"properties"`
}

func (callActivity *TCallActivity) GetID() string {
	return callActivity.ID
}

func (callActivity *TCallActivity) GetName() string {
	return callActivity.Name
}

func (callActivity *TCallActivity) GetType() sepc_element_types.ElementType {
	return sepc_element_types.CallActivity
}

func (callActivity *TCallActivity) GetIncomingAssociation() []string {
	return callActivity.IncomingAssociation
}

func (callActivity *TCallActivity) GetOutgoingAssociation() []string {
	return callActivity.OutgoingAssociation
}
