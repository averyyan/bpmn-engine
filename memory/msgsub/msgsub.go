package msgsub

import (
	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	"github.com/segmentio/ksuid"
)

func New(pi engine_types.ProcessInstance, variables map[string]any, baseElement sepc_types.Message) *MessageSubscription {
	return &MessageSubscription{
		messageID:          baseElement.GetID(),
		messageName:        baseElement.GetName(),
		key:                ksuid.New().String(),
		processInstanceKey: pi.GetKey(),
		consumed:           false,
		variables:          variables,
	}
}

// msgsub *MessageSubscription github.com/averyyan/bpmn-engine/bpmn/engine/types.MessageSubscription
type MessageSubscription struct {
	messageID          string
	messageName        string
	key                string
	processInstanceKey string
	consumed           bool
	variables          map[string]any
}

func (msgsub *MessageSubscription) GetMessageID() string {
	return msgsub.messageID
}

func (msgsub *MessageSubscription) GetMessageName() string {
	return msgsub.messageName
}

func (msgsub *MessageSubscription) GetKey() string {
	return msgsub.key
}

func (msgsub *MessageSubscription) GetProcessInstanceKey() string {
	return msgsub.processInstanceKey
}

func (msgsub *MessageSubscription) GetConsumed() bool {
	return msgsub.consumed
}

func (msgsub *MessageSubscription) GetVariables() map[string]any {
	return msgsub.variables
}

func (msgsub *MessageSubscription) SetConsumed(consumed bool) {
	msgsub.consumed = consumed
}
