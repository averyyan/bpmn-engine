package engine_types

import sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"

// 返回错误代表流程实例失败
type TaskHandler func(acitv ActivatedActivity) error

type TaskHandlerManager interface {
	// 注册 Servcie Task 元素处理函数
	RegisterServiceTaskHandler(taskType string, fun TaskHandler)
	// 获取 Servcie Task 元素处理函数
	GetServiceTaskHandler(taskType string) (TaskHandler, error)
}

type ActivatedActivity interface {
	// 获取流程实例
	GetProcessInstance() ProcessInstance
	// 获取元素
	GetElement() sepc_types.BaseElement
	// 设置活动完成
	Complete()
	// 设置活动失败
	Fail(reason string)
}
