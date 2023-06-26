package event

import sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"

// ted *TTimerEventDefinition github.com/averyyan/bpmn-engine/bpmn/sepc/types.TimerEventDefinition
type TTimerEventDefinition struct {
	ID           string         `xml:"id,attr" json:"id"`                 // 时间中间件ID
	TimeDuration *TTimeDuration `xml:"timeDuration" json:"time_duration"` // 时间中间件延迟时间 ISO 8601 格式 PT15S
	TimeDate     *TTimeDate     `xml:"timeDate" json:"time_date"`         // 时间中间件执行 ISO 8601 格式 2023-10-10T00:00:00Z
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
	XMLText string `xml:",innerxml" json:"text"`
}

func (td *TTimeDuration) GetText() string {
	return td.XMLText
}

type TTimeDate struct {
	XMLText string `xml:",innerxml" json:"text"`
}

func (td *TTimeDate) GetText() string {
	return td.XMLText
}
