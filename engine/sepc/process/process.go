package process

import "github.com/averyyan/bpmn-engine/engine/sepc/event"

// 流程内容
type TProcess struct {
	ID           string `xml:"id,attr"`           // ID
	Name         string `xml:"name,attr"`         // 名称
	IsExecutable bool   `xml:"isExecutable,attr"` // 是否启用
	// 随着流程元素增加添加
	StartEvents             []*event.TStartEvent             `xml:"startEvent"`             // 开始事件
	EndEvents               []*event.TEndEvent               `xml:"endEvent"`               // 结束事件
	IntermediateCatchEvents []*event.TIntermediateCatchEvent `xml:"intermediateCatchEvent"` // 中间事件

}
