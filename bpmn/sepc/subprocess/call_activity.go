package subprocess

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 可重复流程
// callActivity *TCallActivity github.com/averyyan/bpmn-engine/bpmn/sepc/types.BaseElement
// callActivity *TCallActivity github.com/averyyan/bpmn-engine/bpmn/sepc/types.CallActivity
type TCallActivity struct {
	ID                  string                     `xml:"id,attr"`                               // 元素ID
	Name                string                     `xml:"name,attr"`                             // 元素名称
	Documentation       string                     `xml:"documentation,attr"`                    // 元素说明
	IncomingAssociation []string                   `xml:"incoming"`                              // 元素入Flow元素IDs
	OutgoingAssociation []string                   `xml:"outgoing"`                              // 元素出Flow元素IDs
	Properties          []*extensions.TPropertie   `xml:"extensionElements>properties>property"` // 外部数据
	CalledElement       *extensions.TCalledElement `xml:"extensionElements>calledElement"`       // 子流程ID
}

func (callActivity *TCallActivity) HasCalledElement() bool {
	return callActivity.CalledElement != nil
}

func (callActivity *TCallActivity) GetCalledElement() sepc_types.CalledElement {
	return callActivity.CalledElement
}

func (callActivity *TCallActivity) GetDocumentation() string {
	return callActivity.Documentation
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

func (callActivity *TCallActivity) GetProperties() []sepc_types.Propertie {
	return extensions.Properties2Interface(callActivity.Properties)
}
