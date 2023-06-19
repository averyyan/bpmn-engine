package engine_types

import (
	"context"

	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

type ActivityManager interface {
	// 创建新的活动
	Create(ctx context.Context, pi ProcessInstance, baseElement sepc_types.BaseElement) (Activity, error)
	// 按状态查找活动
	FindByStates(ctx context.Context, pi ProcessInstance, states []sepc_element_types.LifecycleState) ([]Activity, error)
	// 设置活动状态为完成
	SetCompleted(ctx context.Context, activ Activity) error
	// 设置活动状态为激活
	SetActive(ctx context.Context, activ Activity) error
	// 设置活动状态为失败
	SetFailed(ctx context.Context, activ Activity, reason string) error
}

type Activity interface {
	BaseElement
}

type ActivatedActivity interface {
	// 获取流程实例
	GetProcessInstance() ProcessInstance
	// 获取活动实例
	GetActivity() Activity
	// 获取元素
	GetElement() sepc_types.ServiceTaskElement
	// 设置活动完成
	Complete() error
	// 设置活动失败
	Fail(reason string) error
}
