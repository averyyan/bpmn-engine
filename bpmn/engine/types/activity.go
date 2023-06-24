package engine_types

import (
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

type ActivityManager interface {
}

type Activity interface {
}

type ActivatedActivity interface {
	// 获取流程实例
	GetProcessInstance() ProcessInstance
	// 获取活动实例
	GetActivity() Activity
	// 获取元素
	GetElement() sepc_types.ServiceTask
	// 设置活动完成
	Complete() error
	// 设置活动失败
	Fail(reason string) error
}
