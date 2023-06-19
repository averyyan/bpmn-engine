package sepc_types

// 任务元素接口
type ExclusiveGateway interface {
	// 基础元素接口
	BaseElement
	// 获取默认序列流
	GetDefault() string
}
