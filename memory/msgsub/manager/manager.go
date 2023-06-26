package memory_msgsub_manager

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	"github.com/averyyan/bpmn-engine/memory/msgsub"
)

func New() *MessageSubscriptionManager {
	return &MessageSubscriptionManager{
		Msgsubs: make(map[string]*msgsub.MessageSubscription),
	}
}

// manager *MessageSubscriptionManager github.com/averyyan/bpmn-engine/bpmn/engine/types.MessageSubscriptionManager
type MessageSubscriptionManager struct {
	Msgsubs map[string]*msgsub.MessageSubscription `json:"msg_subs"`
}

// 通过消息ID找到消息订阅
func (manager *MessageSubscriptionManager) FindOneByID(ctx context.Context, msgID string) (engine_types.MessageSubscription, error) {
	for _, sub := range manager.Msgsubs {
		if sub.GetMessageID() == msgID {
			return sub, nil
		}
	}
	return nil, fmt.Errorf("未找到消息ID【%s】的消息订阅", msgID)
}

// 创建消费订阅
func (manager *MessageSubscriptionManager) Create(ctx context.Context, msgElement sepc_types.Message) (engine_types.MessageSubscription, error) {
	msgsub := msgsub.New(msgElement)
	manager.Msgsubs[msgsub.GetKey()] = msgsub
	return msgsub, nil
}

// 消费消息订阅
func (manager *MessageSubscriptionManager) Consumed(ctx context.Context, msgKey string) error {
	delete(manager.Msgsubs, msgKey)
	return nil
}

// func (manager *MessageSubscriptionManager) FindOneNotConsumedMsgSubByMsgID(ctx context.Context, pi engine_types.ProcessInstance, msgID string) (engine_types.MessageSubscription, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (manager *MessageSubscriptionManager) HasActiveMessageSubscriptionForId(ctx context.Context, pi engine_types.ProcessInstance, id string) (bool, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (manager *MessageSubscriptionManager) SetConsumed(ctx context.Context, pi engine_types.ProcessInstance, msgsub engine_types.MessageSubscription, consumed bool) error {
// 	panic("not implemented") // TODO: Implement
// }

// func (manager *MessageSubscriptionManager) FindOneNotConsumedMsgSubByMsgID(ctx context.Context, pi engine_types.ProcessInstance, msgID string) (engine_types.MessageSubscription, error) {
// 	for _, msgsub := range manager.Msgsubs {
// 		if msgsub.GetProcessInstanceKey() == pi.GetKey() && !msgsub.GetConsumed() && msgsub.GetMessageID() == msgID {
// 			return msgsub, nil
// 		}
// 	}
// 	return nil, fmt.Errorf("未找到msgsub")
// }

// func (manager *MessageSubscriptionManager) HasActiveMessageSubscriptionForId(ctx context.Context, pi engine_types.ProcessInstance, id string) (bool, error) {
// 	panic("not implemented") // TODO: Implement
// }

// func (manager *MessageSubscriptionManager) SetConsumed(ctx context.Context, pi engine_types.ProcessInstance, msgsub engine_types.MessageSubscription, consumed bool) error {
// 	_msgsub := manager.Msgsubs[msgsub.GetKey()]
// 	_msgsub.SetConsumed(consumed)
// 	return nil
// }
