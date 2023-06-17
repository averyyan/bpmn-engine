package sepc_element_types

type ElementType string

const (
	StartEvent             ElementType = "START_EVENT"              // 开始事件
	EndEvent               ElementType = "END_EVENT"                // 结束事件
	ServiceTask            ElementType = "SERVICE_TASK"             // 自动任务
	UserTask               ElementType = "USER_TASK"                // 人工任务
	ParallelGateway        ElementType = "PARALLEL_GATEWAY"         // 并行网关
	ExclusiveGateway       ElementType = "EXCLUSIVE_GATEWAY"        // 逻辑网关
	IntermediateCatchEvent ElementType = "INTERMEDIATE_CATCH_EVENT" // 消息捕获事件
	EventBasedGateway      ElementType = "EVENT_BASED_GATEWAY"      // 基于事件的网关
	SequenceFlow           ElementType = "SEQUENCE_FLOW"            // 箭头
	CallActivity           ElementType = "CALLACTIVITY"             // 重复子流程
)
