package event

import (
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// ice *TIntermediateCatchEvent github.com/averyyan/bpmn-engine/bpmn/sepc/types.IntermediateCatchEvent
type TIntermediateCatchEvent struct {
	ID                     string                   `xml:"id,attr"`                // 元素ID
	Name                   string                   `xml:"name,attr"`              // 元素名称
	IncomingAssociation    []string                 `xml:"incoming"`               // 元素入Flow元素IDs
	OutgoingAssociation    []string                 `xml:"outgoing"`               // 元素出Flow元素IDs
	MessageEventDefinition *TMessageEventDefinition `xml:"messageEventDefinition"` // 基于消息中间件内容
	TimerEventDefinition   *TTimerEventDefinition   `xml:"timerEventDefinition"`   // 基于时间中间件内容
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
	return sepc_element_types.IntermediateCatchEvent
}

func (ice *TIntermediateCatchEvent) GetMessageEventDefinition() sepc_types.MessageEventDefinition {
	return ice.MessageEventDefinition
}

func (ice *TIntermediateCatchEvent) GetTimerEventDefinition() sepc_types.TimerEventDefinition {
	return ice.TimerEventDefinition
}

type TMessageEventDefinition struct {
	ID         string `xml:"id,attr"`         // 消息中间件ID
	MessageRef string `xml:"messageRef,attr"` // 消息中间件匹配消息订阅ID
}

func (med *TMessageEventDefinition) GetID() string {
	return med.ID
}

func (med *TMessageEventDefinition) GetMessageRef() string {
	return med.MessageRef
}

type TTimerEventDefinition struct {
	ID           string         `xml:"id,attr"`      // 时间中间件ID
	TimeDuration *TTimeDuration `xml:"timeDuration"` // 时间中间件延迟时间 ISO 8601 格式 PT15S
	TimeDate     *TTimeDate     `xml:"timeDate"`     // 时间中间件执行 ISO 8601 格式 2023-10-10T00:00:00Z
}

func (ted *TTimerEventDefinition) GetID() string {
	return ted.ID
}

func (ted *TTimerEventDefinition) GetTimeDuration() sepc_types.BaseElementText {
	return ted.TimeDuration
}

func (ted *TTimerEventDefinition) GetTimeDate() sepc_types.BaseElementText {
	return ted.TimeDate
}

type TTimeDuration struct {
	XMLText string `xml:",innerxml"`
}

func (td *TTimeDuration) GetText() string {
	return td.XMLText
}

type TTimeDate struct {
	XMLText string `xml:",innerxml"`
}

func (td *TTimeDate) GetText() string {
	return td.XMLText
}
