package engine

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/event"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 中间事件元素处理
func handleMessageIntermediateCatchEvent(
	ctx context.Context,
	pi engine_types.ProcessInstance,
	iceElement *event.TIntermediateCatchEvent,
) (bool, error) {
	var ice engine_types.BaseElement
	var err error
	// 找到或者创建激活的消息
	if ice, err = pi.GetElementManager().FindOneMessageICE(ctx, iceElement.ID); err != nil {
		ice, err = pi.GetElementManager().CreateMessageICE(ctx, iceElement)
		if err != nil {
			return false, err
		}
	}
	// 通过ice信息找到订阅消息
	msgsub := findMatchingMessageSubscription(ctx, pi, iceElement)
	// 如果存在对应的消息订阅则将订阅消费掉，并将ice设置为完成状态
	if msgsub != nil {
		if err := pi.GetMessageSubscriptionManager().Consumed(ctx, msgsub.GetKey()); err != nil {
			return false, err
		}
		if err := pi.GetElementManager().DeleteMessageICE(ctx, ice.GetKey()); err != nil {
			return false, err
		}
		return true, nil
	}
	return false, nil
}

// 找到消息元素对应的消息订阅
func findMatchingMessageSubscription(
	ctx context.Context,
	pi engine_types.ProcessInstance,
	iceElement sepc_types.IntermediateCatchEvent,
) engine_types.MessageSubscription {
	msg, err := findMessageById(pi, iceElement.GetMessageEventDefinition().GetMessageRef())
	if err != nil {
		return nil
	}
	msgsub, err := pi.GetMessageSubscriptionManager().FindOneByID(ctx, msg.GetID())
	if err != nil {
		return nil
	}
	return msgsub
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
	pi engine_types.ProcessInstance,
) error {
	// 检查未完成元素
	var ices []sepc_types.BaseElement
	for _, element := range pi.GetElementManager().FindActiveElements(ctx) {
		switch element.GetType() {
		case sepc_element_types.CallActivity, sepc_element_types.UserTask: // 存在未完成的任务
			return nil
		case sepc_element_types.MessageIntermediateCatchEvent:
			ices = append(ices, element)
		}
	}
	// 检查未完成的事件
	hasActiveIce, err := checkUnCompletedICE(ctx, pi, ices)
	if err != nil {
		return err
	}
	if hasActiveIce {
		return nil
	}
	// 流程已经完结
	if err := pi.SetCompleted(ctx); err != nil {
		return err
	}
	// 运行父级流程
	if len(pi.GetParentProcessInstanceKey()) > 0 {
		if _, err := ContinuePIByKey(ctx, pi.GetParentProcessInstanceKey()); err != nil {
			return err
		}
	}
	return nil
}

func checkUnCompletedICE(
	ctx context.Context,
	pi engine_types.ProcessInstance,
	iceElements []sepc_types.BaseElement,
) (bool, error) {
	definitions, err := pi.GetDefinitions()
	if err != nil {
		return false, err
	}
	activedICE := make(map[string]bool)
	for _, ice := range iceElements {
		activedICE[ice.GetID()] = true
	}
	if err != nil {
		return false, err
	}
	for _, gateway := range definitions.Process.EventBasedGateways {
		flows := findSequenceFlows(definitions.Process.SequenceFlows, gateway.OutgoingAssociation)
		isOneEventCompleted := true
		for _, flow := range flows {
			isOneEventCompleted = isOneEventCompleted && !activedICE[flow.TargetRef]
		}
		for _, flow := range flows {
			activedICE[flow.TargetRef] = isOneEventCompleted
		}
	}
	for _, v := range activedICE {
		if v {
			return true, nil
		}
	}
	return false, nil
}
