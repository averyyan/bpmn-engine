package activity

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 可重复流程
// callActivity *TCallActivity github.com/averyyan/bpmn-engine/bpmn/sepc/types.BaseElement
// callActivity *TCallActivity github.com/averyyan/bpmn-engine/bpmn/sepc/types.CallActivity
type TCallActivity struct {
	ID                  string                     `xml:"id,attr" json:"id"`                                       // 元素ID
	Name                string                     `xml:"name,attr" json:"name"`                                   // 元素名称
	Documentation       string                     `xml:"documentation,attr" json:"documentation"`                 // 元素说明
	IncomingAssociation []string                   `xml:"incoming" json:"incoming"`                                // 元素入Flow元素IDs
	OutgoingAssociation []string                   `xml:"outgoing" json:"outgoing"`                                // 元素出Flow元素IDs
	Properties          []*extensions.TPropertie   `xml:"extensionElements>properties>property" json:"properties"` // 外部数据
	CalledElement       *extensions.TCalledElement `xml:"extensionElements>calledElement" json:"called_element"`   // 子流程ID
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

func (callActivity *TCallActivity) GetDocumentation() string {
	return callActivity.Documentation
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

func (callActivity *TCallActivity) HasCalledElement() bool {
	return callActivity.CalledElement != nil
}

func (callActivity *TCallActivity) GetCalledElement() sepc_types.CalledElement {
	return callActivity.CalledElement
}

// // 判断是否存在重复的流程ID
// func (callActivity *TCallActivity) HasCalledElement() bool {
// 	return callActivity.CalledElement != nil && callActivity.CalledElement.ProcessID != ""
// }

// // 获取流程信息内容
// func (callActivity *TCallActivity) GetCalledElement() sepc_types.CalledElement {
// 	return callActivity.CalledElement
// }

// func (callActivity *TCallActivity) GetDocumentation() string {
// 	return callActivity.Documentation
// }

// func (callActivity *TCallActivity) GetID() string {
// 	return callActivity.ID
// }

// func (callActivity *TCallActivity) GetName() string {
// 	return callActivity.Name
// }

// func (callActivity *TCallActivity) GetType() sepc_element_types.ElementType {
// 	return sepc_element_types.CallActivity
// }

// func (callActivity *TCallActivity) GetIncomingAssociation() []string {
// 	return callActivity.IncomingAssociation
// }

// func (callActivity *TCallActivity) GetOutgoingAssociation() []string {
// 	return callActivity.OutgoingAssociation
// }

// func (callActivity *TCallActivity) GetProperties() []sepc_types.Propertie {
// 	return extensions.Properties2Interface(callActivity.Properties)
// }
