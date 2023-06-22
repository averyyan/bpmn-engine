package engine

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

func handleIntermediateCatchEvent(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
	ice sepc_types.IntermediateCatchEvent,
) (bool, error) {
	if ice.GetMessageEventDefinition() != nil {
		return handleIntermediateMessageCatchEvent(ctx, state, pi, ice)
	}
	return false, fmt.Errorf("未支持此类型事件")
}

// 中间事件元素处理
func handleIntermediateMessageCatchEvent(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
	iceElement sepc_types.IntermediateCatchEvent,
) (bool, error) {
	var ice engine_types.IntermediateCatchEvent
	var err error
	// 找到或者创建激活的消息
	if ice, err = state.ICEManager().FindOneMsgICEByStateAndID(
		ctx, pi, sepc_element_types.Active, iceElement.GetID(),
	); err != nil {
		// 未存在则创建新消息中间事件
		if ice, err = state.ICEManager().CreateMsgICE(ctx, pi, iceElement); err != nil {
			return false, err
		}
	}
	// 通过ice信息找到订阅消息
	msgsub, _ := findMatchingMessageSubscription(ctx, state, pi, iceElement)
	// 如果存在对应的消息订阅则将订阅消费掉，并将ice设置为完成状态
	if msgsub != nil {
		if err := state.MessageSubscriptionManager().SetConsumed(ctx, pi, msgsub, true); err != nil {
			return false, err
		}
		if err := state.ICEManager().SetCompleted(ctx, ice); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// 找到消息元素对应的消息订阅
func findMatchingMessageSubscription(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
	ice sepc_types.IntermediateCatchEvent,
) (engine_types.MessageSubscription, error) {
	msg, err := findMessageById(pi, ice.GetMessageEventDefinition().GetMessageRef())
	if err != nil {
		return nil, err
	}
	msgsub, err := state.MessageSubscriptionManager().FindOneNotConsumedMsgSubByMsgID(ctx, pi, msg.GetID())
	if err != nil {
		return nil, err
	}
	return msgsub, nil
}

// 通过消息ID找到消息
func findMessageById(pi engine_types.ProcessInstance, msgID string) (sepc_types.Message, error) {
	definitions, err := pi.GetDefinitions()
	if err != nil {
		return nil, err
	}
	for _, msgElement := range definitions.Messages {
		if msgElement.ID == msgID {
			return msgElement, nil
		}
	}
	return nil, fmt.Errorf("未存在ID【%s】的消息元素", msgID)
}

// 处理EndEvent元素
func handleEndEvent(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
) error {
	// 检查未完成的任务
	activs, err := state.ActivityManager().FindByStates(
		ctx, pi,
		[]sepc_element_types.LifecycleState{
			sepc_element_types.Ready,
			sepc_element_types.Active,
		},
	)
	if err != nil {
		return err
	}
	// 判断流程是否完结
	if len(activs) > 0 {
		return nil
	}
	// 流程已经完结
	if err := state.ProcessInstanceManager().SetCompleted(ctx, pi); err != nil {
		return err
	}
	return nil
}
