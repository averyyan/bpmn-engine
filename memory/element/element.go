package memory_element

import (
	"time"

	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	"github.com/segmentio/ksuid"
)

func New(baseElement sepc_types.BaseElement) BaseElement {
	return BaseElement{
		ElementType: baseElement.GetType(),
		Key:         ksuid.New().String(),
		CreateTime:  time.Now().Local(),
	}
}

// element *BaseElement github.com/averyyan/bpmn-engine/bpmn/engine/types.BaseElement
type BaseElement struct {
	ElementType sepc_element_types.ElementType `json:"element_type"`
	Key         string                         `json:"key"`
	CreateTime  time.Time                      `json:"create_time"`
}

// 获取元素类型
func (element *BaseElement) GetElementType() sepc_element_types.ElementType {
	return element.ElementType
}

// 元素唯一Key
func (element *BaseElement) GetKey() string {
	return element.Key
}
