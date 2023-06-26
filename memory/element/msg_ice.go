package memory_element

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/event"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

func NewMessageICE(element *event.TIntermediateCatchEvent) *MessageICE {
	return &MessageICE{
		Element:     element,
		BaseElement: New(element),
	}
}

type MessageICE struct {
	BaseElement `json:",inline"`
	Element     *event.TIntermediateCatchEvent `json:"element"`
}

// 获取元素
func (element *MessageICE) GetElement() sepc_types.BaseElement {
	return element.Element
}
