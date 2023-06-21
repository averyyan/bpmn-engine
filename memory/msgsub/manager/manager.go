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
		msgsubs: make(map[string]*msgsub.MessageSubscription),
	}
}

// manager *MessageSubscriptionManager github.com/averyyan/bpmn-engine/bpmn/engine/types.MessageSubscriptionManager
type MessageSubscriptionManager struct {
	msgsubs map[string]*msgsub.MessageSubscription
}

func (manager *MessageSubscriptionManager) Create(ctx context.Context, pi engine_types.ProcessInstance, variables map[string]any, msgElement sepc_types.Message) (engine_types.MessageSubscription, error) {
	msgsub := msgsub.New(pi, variables, msgElement)
	manager.msgsubs[msgsub.GetKey()] = msgsub
	return msgsub, nil
}

func (manager *MessageSubscriptionManager) FindOneNotConsumedMsgSubByMsgID(ctx context.Context, pi engine_types.ProcessInstance, msgID string) (engine_types.MessageSubscription, error) {
	for _, msgsub := range manager.msgsubs {
		if msgsub.GetProcessInstanceKey() == pi.GetKey() && !msgsub.GetConsumed() && msgsub.GetMessageID() == msgID {
			return msgsub, nil
		}
	}
	return nil, fmt.Errorf("未找到msgsub")
}

func (manager *MessageSubscriptionManager) HasActiveMessageSubscriptionForId(ctx context.Context, pi engine_types.ProcessInstance, id string) (bool, error) {
	panic("not implemented") // TODO: Implement
}

func (manager *MessageSubscriptionManager) SetConsumed(ctx context.Context, pi engine_types.ProcessInstance, msgsub engine_types.MessageSubscription, consumed bool) error {
	_msgsub := manager.msgsubs[msgsub.GetKey()]
	_msgsub.SetConsumed(consumed)
	return nil
}
