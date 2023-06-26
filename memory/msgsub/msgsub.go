package msgsub

import (
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	"github.com/segmentio/ksuid"
)

func New(baseElement sepc_types.Message) *MessageSubscription {
	return &MessageSubscription{
		MessageID:   baseElement.GetID(),
		MessageName: baseElement.GetName(),
		Key:         ksuid.New().String(),
	}
}

// msgsub *MessageSubscription github.com/averyyan/bpmn-engine/bpmn/engine/types.MessageSubscription
type MessageSubscription struct {
	MessageID   string `json:"message_id"`
	MessageName string `json:"message_name"`
	Key         string `json:"key"`
}

func (msgsub *MessageSubscription) GetMessageID() string {
	return msgsub.MessageID
}

func (msgsub *MessageSubscription) GetMessageName() string {
	return msgsub.MessageName
}

func (msgsub *MessageSubscription) GetKey() string {
	return msgsub.Key
}
