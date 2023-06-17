package sepc_types

type StartEvent interface {
	BaseElement // 基础接口组合
}

type EndEvent interface {
	BaseElement // 基础接口组合
}

type IntermediateCatchEvent interface {
	BaseElement                                        // 基础接口组合
	GetMessageEventDefinition() MessageEventDefinition // 获取基于消息的中间事件详情
	GetTimerEventDefinition() TimerEventDefinition     // 获取基于定时器的中间事件详情
}

type MessageEventDefinition interface {
	GetID() string
	GetMessageRef() string
}

type TimerEventDefinition interface {
	GetID() string
	GetTimeDuration() BaseElementText
	GetTimeDate() BaseElementText
}
