package engine_types

// 流程引擎接口
type Engine interface {
	ProcessInstanceManager() ProcessInstanceManager // 流程实例管理
}
