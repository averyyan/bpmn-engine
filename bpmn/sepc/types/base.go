package sepc_types

import sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"

// 基础元素接口
type BaseElement interface {
	GetID() string                           // 获取元素ID
	GetName() string                         // 获取元素名称
	GetType() sepc_element_types.ElementType // 获取元素类型
	GetIncomingAssociation() []string        // 获取元素入Flow元素IDs
	GetOutgoingAssociation() []string        // 获取元素出Flow元素IDs
}

// 元素获取XML内容接口
type BaseElementText interface {
	GetText() string
}
