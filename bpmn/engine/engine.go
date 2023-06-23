package engine

import (
	"context"
	"fmt"
	"os"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
)

// 根据BPMN文件地址创建流程实例
func CreateInstanceByFileAndRun(
	ctx context.Context,
	state engine_types.Engine,
	filepath string,
	variables map[string]any,
) (engine_types.ProcessInstance, error) {
	raw, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	process, err := state.ProcessManager().Create(ctx, raw)
	if err != nil {
		return nil, err
	}
	pi, err := state.ProcessInstanceManager().Create(ctx, process.GetRaw(), "", variables)
	if err != nil {
		return nil, err
	}
	return pi, run(ctx, state, pi)
}

// 继续流程实例
func ContinueProcessInstance(
	ctx context.Context,
	state engine_types.Engine,
	piKey string,
) (engine_types.ProcessInstance, error) {
	pi, err := state.ProcessInstanceManager().FindOneByKey(ctx, piKey)
	if err != nil {
		return nil, err
	}
	return pi, run(ctx, state, pi)
}

// 创建子流程
func createChildInstance(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
	callActivityElement sepc_types.CallActivity,
) (engine_types.BaseElement, error) {
	if !callActivityElement.HasCalledElement() {
		return nil, fmt.Errorf("未存在子流程信息")
	}
	cactiv, err := state.CallActivityManager().(engine_types.BaseManager).Create(
		ctx, pi, callActivityElement.(sepc_types.BaseElement),
	)
	if err != nil {
		return nil, err
	}
	process, err := state.ProcessManager().FindOneByID(ctx, callActivityElement.GetCalledElement().GetProcessID())
	if err != nil {
		return nil, err
	}
	childPI, err := state.ProcessInstanceManager().Create(ctx, process.GetRaw(), pi.GetKey(), pi.GetVariables())
	if err != nil {
		return nil, err
	}
	if err := state.CallActivityManager().(engine_types.BaseManager).SetActive(ctx, cactiv); err != nil {
		return nil, err
	}
	if err := state.CallActivityManager().SetChildPIKey(ctx, cactiv, childPI.GetKey()); err != nil {
		return nil, err
	}
	return cactiv, nil
}

// 基于消息名称对流程进行重启
func PublishEventForInstanceAndRun(
	ctx context.Context,
	state engine_types.Engine,
	piKey string,
	messageName string,
	variables map[string]any,
) (engine_types.ProcessInstance, error) {
	pi, err := state.ProcessInstanceManager().FindOneByKey(ctx, piKey)
	if err != nil {
		return nil, err
	}
	switch pi.GetState() {
	case sepc_pi_types.Completed:
		return nil, fmt.Errorf("流程Key【%s】已完成", pi.GetKey())
	}
	definitions, err := pi.GetDefinitions()
	if err != nil {
		return nil, err
	}
	// 创建新的消息订阅
	for _, msg := range definitions.Messages {
		if msg.Name == messageName {
			if _, err := state.MessageSubscriptionManager().Create(
				ctx,
				pi,
				variables,
				msg,
			); err != nil {
				return nil, err
			}
			return pi, run(ctx, state, pi)
		}
	}
	return nil, fmt.Errorf("创建消息订阅失败")
}

// 流程运行核心
func run(ctx context.Context, state engine_types.Engine, pi engine_types.ProcessInstance) error {
	definitions, err := pi.GetDefinitions()
	if err != nil {
		return err
	}
	elementMap := makeElementMap(definitions)
	type queueElement struct {
		inboundFlowId string
		baseElement   sepc_types.BaseElement
	}
	queue := make([]queueElement, 0)
	switch pi.GetState() {
	case sepc_pi_types.Ready:
		// 将开始事件排入待处理元素队列
		for _, event := range definitions.Process.StartEvents {
			queue = append(queue, queueElement{
				inboundFlowId: "",
				baseElement:   event,
			})
		}
		// 设置流程实例为激活状态
		if err := state.ProcessInstanceManager().SetActive(ctx, pi); err != nil {
			return err
		}
	case sepc_pi_types.Active:
		// 找到未完成的子流程
		cactivs, err := findCallActivitysForContinuation(ctx, state, pi)
		if err != nil {
			return err
		}
		for _, cactiv := range cactivs {
			queue = append(queue, queueElement{
				inboundFlowId: "",
				baseElement:   cactiv,
			})
		}
		// 找到未完成的中间事件
		intermediateCatchEvents, err := findIntermediateCatchEventsForContinuation(ctx, state, pi)
		if err != nil {
			return err
		}
		for _, ice := range intermediateCatchEvents {
			queue = append(queue, queueElement{
				inboundFlowId: "",
				baseElement:   ice,
			})
		}
	case sepc_pi_types.Completed:
	case sepc_pi_types.Failed:

	default:
		return fmt.Errorf("未存在此流程实例状态")
	}
	for len(queue) > 0 {
		element := queue[0].baseElement
		inboundFlowId := queue[0].inboundFlowId
		queue = queue[1:]
		continueNextElement, err := handleElement(ctx, state, pi, element)
		if err != nil {
			return err
		}
		if continueNextElement {
			if inboundFlowId != "" {
				if err := state.ProcessInstanceManager().RemoveScheduledFlows(ctx, pi, inboundFlowId); err != nil {
					return err
				}
			}
			nextFlows := findSequenceFlows(definitions.Process.SequenceFlows, element.GetOutgoingAssociation())
			if element.GetType() == sepc_element_types.ExclusiveGateway {
				_nextFlows, err := exclusivelyFilterByConditionExpression(nextFlows, pi.GetVariables(), element.(sepc_types.ExclusiveGateway))
				if err != nil {
					return err
				}
				nextFlows = _nextFlows
			}
			for _, nextFlow := range nextFlows {
				if err := state.ProcessInstanceManager().AppendScheduledFlows(ctx, pi, nextFlow.ID); err != nil {
					return err
				}
				if baseElement, ok := elementMap[nextFlow.TargetRef]; ok {
					queue = append(queue, queueElement{
						inboundFlowId: nextFlow.ID,
						baseElement:   baseElement,
					})
				} else {
					return fmt.Errorf("元素【%s】未找到", nextFlow.ID)
				}
			}
		}
	}
	return nil
}

func findCallActivitysForContinuation(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
) ([]sepc_types.BaseElement, error) {
	var baseElements []sepc_types.BaseElement
	definitions, err := pi.GetDefinitions()
	if err != nil {
		return nil, err
	}
	cacitvs, err := state.CallActivityManager().(engine_types.BaseManager).FindByStates(
		ctx, pi,
		[]sepc_element_types.LifecycleState{sepc_element_types.Ready, sepc_element_types.Active},
	)
	if err != nil {
		return nil, err
	}
	for _, cacitv := range cacitvs {
		for _, callactivityElement := range definitions.Process.CallActivitys {
			if callactivityElement.GetID() == cacitv.GetElementID() {
				baseElements = append(baseElements, callactivityElement)
			}
		}
	}
	return baseElements, nil
}

func findIntermediateCatchEventsForContinuation(ctx context.Context, state engine_types.Engine, pi engine_types.ProcessInstance) ([]sepc_types.BaseElement, error) {
	var events []sepc_types.BaseElement
	definitions, err := pi.GetDefinitions()
	if err != nil {
		return nil, err
	}
	msgIces, err := state.ICEManager().(engine_types.BaseManager).FindByStates(
		ctx, pi,
		[]sepc_element_types.LifecycleState{sepc_element_types.Active},
	)
	if err != nil {
		return nil, err
	}
	for _, ice := range definitions.Process.IntermediateCatchEvents {
		for _, msgIce := range msgIces {
			if ice.ID == msgIce.GetElementID() {
				msgID := ice.GetMessageEventDefinition().GetMessageRef()
				for _, message := range definitions.Messages {
					if message.ID == msgID {
						events = append(events, ice)
					}
				}
			}
		}
	}
	return events, nil
}
