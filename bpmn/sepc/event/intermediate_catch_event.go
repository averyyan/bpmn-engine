package event

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// ice *TIntermediateCatchEvent github.com/averyyan/bpmn-engine/bpmn/sepc/types.BaseElement
// ice *TIntermediateCatchEvent github.com/averyyan/bpmn-engine/bpmn/sepc/types.IntermediateCatchEvent
type TIntermediateCatchEvent struct {
	ID                     string                   `xml:"id,attr" json:"id"`                                       // 元素ID
	Name                   string                   `xml:"name,attr" json:"name"`                                   // 元素名称
	Documentation          string                   `xml:"documentation,attr" json:"documentation"`                 // 元素说明
	IncomingAssociation    []string                 `xml:"incoming" json:"incoming"`                                // 元素入Flow元素IDs
	OutgoingAssociation    []string                 `xml:"outgoing" json:"outgoing"`                                // 元素出Flow元素IDs
	Properties             []*extensions.TPropertie `xml:"extensionElements>properties>property" json:"properties"` // 外部数据
	MessageEventDefinition *TMessageEventDefinition `xml:"messageEventDefinition" json:"message_event_definition"`  // 基于消息中间件内容
	TimerEventDefinition   *TTimerEventDefinition   `xml:"timerEventDefinition" json:"timer_event_Definition"`      // 基于时间中间件内容
}

func (ice *TIntermediateCatchEvent) GetDocumentation() string {
	return ice.Documentation
}

func (ice *TIntermediateCatchEvent) GetProperties() []sepc_types.Propertie {
	return extensions.Properties2Interface(ice.Properties)
}

func (ice *TIntermediateCatchEvent) GetID() string {
	return ice.ID
}

func (ice *TIntermediateCatchEvent) GetName() string {
	return ice.Name
}

func (ice *TIntermediateCatchEvent) GetIncomingAssociation() []string {
	return ice.IncomingAssociation
}

func (ice *TIntermediateCatchEvent) GetOutgoingAssociation() []string {
	return ice.OutgoingAssociation
}

func (ice *TIntermediateCatchEvent) GetType() sepc_element_types.ElementType {
	if ice.HasMessageEventDefinition() {
		return sepc_element_types.MessageIntermediateCatchEvent
	}
	if ice.HasTimerEventDefinition() {
		return sepc_element_types.TimerIntermediateCatchEvent
	}
	return sepc_element_types.IntermediateCatchEvent
}

// 是否存在消息信息
func (ice *TIntermediateCatchEvent) HasMessageEventDefinition() bool {
	return ice.MessageEventDefinition != nil
}

// 是否存在时间
func (ice *TIntermediateCatchEvent) HasTimerEventDefinition() bool {
	return ice.TimerEventDefinition != nil
}

func (ice *TIntermediateCatchEvent) GetMessageEventDefinition() sepc_types.MessageEventDefinition {
	return ice.MessageEventDefinition
}

func (ice *TIntermediateCatchEvent) GetTimerEventDefinition() sepc_types.TimerEventDefinition {
	return ice.TimerEventDefinition
}
