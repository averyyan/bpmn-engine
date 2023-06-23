package engine_types

import (
	"context"
	"time"

	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

type BaseManager interface {
	// 创建新的活动
	Create(ctx context.Context, pi ProcessInstance, baseElement sepc_types.BaseElement) (BaseElement, error)
	// 按状态查找活动
	FindByStates(ctx context.Context, pi ProcessInstance, states []sepc_element_types.LifecycleState) ([]BaseElement, error)
	// 按元素ID和元素状态查找元素
	FindOneByStateAndID(ctx context.Context, pi ProcessInstance, state sepc_element_types.LifecycleState, elementID string) (BaseElement, error)
	// 设置活动状态为完成
	SetCompleted(ctx context.Context, base BaseElement) error
	// 设置活动状态为激活
	SetActive(ctx context.Context, base BaseElement) error
	// 设置活动状态为失败
	SetFailed(ctx context.Context, base BaseElement, reason string) error
}

type BaseElement interface {
	GetElementID() string
	GetElementName() string
	GetKey() string
	GetProcessInstanceKey() string
	GetCreateTime() time.Time
	GetState() sepc_element_types.LifecycleState
}
