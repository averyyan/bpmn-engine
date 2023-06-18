package engine_types

// 流程引擎接口
type Engine interface {
	ProcessInstanceManager() ProcessInstanceManager // 流程实例管理
	ActivityManager() ActivityManager               // 活动管理
	TaskHandlerManager() TaskHandlerManager         // 服务任务管理
}
