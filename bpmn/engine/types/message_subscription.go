package engine_types

import (
	"context"

	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

type MessageSubscriptionManager interface {
	// 创建一个新的消息订阅
	Create(ctx context.Context, msgElement sepc_types.Message) (MessageSubscription, error)
	// 消费消息订阅
	Consumed(ctx context.Context, msgKey string) error
	// 通过消息ID找到消息订阅
	FindOneByID(ctx context.Context, msgID string) (MessageSubscription, error)
}

type MessageSubscription interface {
	GetMessageID() string
	GetMessageName() string
	GetKey() string
}
