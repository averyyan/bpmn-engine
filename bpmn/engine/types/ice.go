package engine_types

import (
	"context"

	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 中间事件管理器
type ICEManager interface {
	// 创建消息中间事件
	CreateMsgICE(ctx context.Context, pi ProcessInstance, baseElement sepc_types.IntermediateCatchEvent) (IntermediateCatchEvent, error)
	// 按状态查找消息中间事件
	FindMsgICEByStates(ctx context.Context, pi ProcessInstance, states []sepc_element_types.LifecycleState) ([]IntermediateCatchEvent, error)
	// 按照状态和ID找到消息中间事件
	FindOneMsgICEByStateAndID(ctx context.Context, pi ProcessInstance, state sepc_element_types.LifecycleState, elementID string) (IntermediateCatchEvent, error)
	// 设置活动状态为完成
	SetCompleted(ctx context.Context, ice IntermediateCatchEvent) error
	// 设置活动状态为失败
	SetFailed(ctx context.Context, ice IntermediateCatchEvent) error
}

type IntermediateCatchEvent interface {
	BaseElement
}
