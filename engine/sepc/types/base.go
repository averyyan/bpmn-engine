package sepc_types

// 基础元素接口
type BaseElement interface {
	GetID() string                    // 获取元素ID
	GetName() string                  // 获取元素名称
	GetIncomingAssociation() []string // 获取元素入Flow元素IDs
	GetOutgoingAssociation() []string // 获取元素出Flow元素IDs
}

// 元素获取XML内容接口
type BaseElementText interface {
	GetText() string
}
