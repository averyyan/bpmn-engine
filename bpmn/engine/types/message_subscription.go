package engine_types

import (
	"context"

	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

type MessageSubscriptionManager interface {
	// 创建一个新的消息订阅
	Create(ctx context.Context, pi ProcessInstance, variables map[string]any, msgElement sepc_types.Message) (MessageSubscription, error) // 创建新的消息订阅
	FindOneNotConsumedMsgSubByMsgID(ctx context.Context, pi ProcessInstance, msgID string) (MessageSubscription, error)                   // 找到未被消耗的消息
	HasActiveMessageSubscriptionForId(ctx context.Context, pi ProcessInstance, id string) (bool, error)                                   // 是否存在激活的信息订阅
	SetConsumed(ctx context.Context, pi ProcessInstance, msgsub MessageSubscription, consumed bool) error
}

type MessageSubscription interface {
	GetMessageID() string
	GetMessageName() string
	GetKey() string
	GetProcessInstanceKey() string
	GetConsumed() bool
	GetVariables() map[string]any
}
