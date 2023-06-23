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
	if ice, err = state.ICEManager().(engine_types.BaseManager).FindOneByStateAndID(
		ctx, pi, sepc_element_types.Active, iceElement.(sepc_types.BaseElement).GetID(),
	); err != nil {
		// 未存在则创建新消息中间事件
		if ice, err = state.ICEManager().(engine_types.BaseManager).Create(ctx, pi, iceElement.(sepc_types.BaseElement)); err != nil {
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
		if err := state.ICEManager().(engine_types.BaseManager).SetCompleted(ctx, ice.(engine_types.BaseElement)); err != nil {
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
	iceElement sepc_types.IntermediateCatchEvent,
) (engine_types.MessageSubscription, error) {
	msg, err := findMessageById(pi, iceElement.GetMessageEventDefinition().GetMessageRef())
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
	hasActivity, err := checkUnCompleted(ctx, state.ActivityManager().(engine_types.BaseManager), pi)
	if err != nil {
		return err
	}
	if hasActivity {
		return nil
	}
	// 检查未完成的子流程
	hasCallActivity, err := checkUnCompleted(ctx, state.CallActivityManager().(engine_types.BaseManager), pi)
	if err != nil {
		return err
	}
	if hasCallActivity {
		return nil
	}
	// 检查未完成的事件
	hasActiveIce, err := checkUnCompletedICE(ctx, state, pi)
	if err != nil {
		return err
	}
	if hasActiveIce {
		return nil
	}
	// 流程已经完结
	if err := state.ProcessInstanceManager().SetCompleted(ctx, pi); err != nil {
		return err
	}
	// 运行父级流程
	if len(pi.GetParentProcessInstanceKey()) > 0 {
		if _, err := ContinueProcessInstance(ctx, state, pi.GetParentProcessInstanceKey()); err != nil {
			return err
		}
	}
	return nil
}

func checkUnCompletedICE(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
) (bool, error) {
	definitions, err := pi.GetDefinitions()
	if err != nil {
		return false, err
	}
	activedICE := make(map[string]bool)
	ices, err := state.ICEManager().(engine_types.BaseManager).FindByStates(
		ctx,
		pi,
		[]sepc_element_types.LifecycleState{
			sepc_element_types.Active,
			sepc_element_types.Ready,
			sepc_element_types.Completed,
		},
	)

	for _, ice := range ices {
		activedICE[ice.GetElementID()] = (ice.GetState() == sepc_element_types.Ready || ice.GetState() == sepc_element_types.Active)
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

func checkUnCompleted(
	ctx context.Context,
	manager engine_types.BaseManager,
	pi engine_types.ProcessInstance,
) (bool, error) {
	datas, err := manager.FindByStates(
		ctx, pi,
		[]sepc_element_types.LifecycleState{
			sepc_element_types.Ready,
			sepc_element_types.Active,
		},
	)
	if err != nil {
		return false, err
	}
	return len(datas) > 0, nil
}
