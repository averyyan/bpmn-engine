package sepc_element_types

// 活动的生命周期 pdf 429页
type LifecycleState string

const (
	Ready        LifecycleState = "READY"        // 准备状态
	Active       LifecycleState = "ACTIVE"       // 活动状态
	WithDrawn    LifecycleState = "WITHDRAWN"    // 撤回(未启用)
	Completing   LifecycleState = "COMPLETING"   // 完成中(未启用)
	Completed    LifecycleState = "COMPLETED"    // 已完成
	Failing      LifecycleState = "FAILING"      // 失败中(未启用)
	Terminating  LifecycleState = "TERMINATING"  // 停止中(未启用)
	Compensating LifecycleState = "COMPENSATING" // 修正中(未启用)
	Failed       LifecycleState = "FAILED"       // 已失败
	Terminated   LifecycleState = "TERMINATED"   // 已停止(未启用)
	Compensated  LifecycleState = "COMPENSATED"  // 已修正
)

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
