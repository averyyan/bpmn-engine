package engine_types

// 流程引擎接口
type Engine interface {
	ProcessManager() ProcessManager                         // 流程管理
	ProcessInstanceManager() ProcessInstanceManager         // 流程实例管理
	ActivityManager() ActivityManager                       // 活动管理
	TaskHandlerManager() TaskHandlerManager                 // 服务任务管理
	ICEManager() ICEManager                                 // 中间事件管理
	MessageSubscriptionManager() MessageSubscriptionManager // 消息注册
	CallActivityManager() CallActivityManager               // 子流程管理
}
