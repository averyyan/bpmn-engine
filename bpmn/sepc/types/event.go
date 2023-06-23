package sepc_types

// 中间事件
type IntermediateCatchEvent interface {
	GetMessageEventDefinition() MessageEventDefinition // 获取基于消息的中间事件详情
	GetTimerEventDefinition() TimerEventDefinition     // 获取基于定时器的中间事件详情
}

// 基于消息的中间事件
type MessageEventDefinition interface {
	GetID() string         // 获取消息中间事件ID
	GetMessageRef() string // 获取消息ID
}

// 基于定时器的中间事件 (两种类型)
type TimerEventDefinition interface {
	GetID() string                    // 获取定时器中间事件ID
	GetTimeDuration() BaseElementText // 定时器延时时间
	GetTimeDate() BaseElementText     // 定时器执行时间
}
