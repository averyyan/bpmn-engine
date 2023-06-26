package engine

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
)

func CreatePIByIDAndRun(ctx context.Context, processID string, variables map[string]any) (engine_types.ProcessInstance, error) {
	pi, err := CreatePIByID(processID, variables)
	if err != nil {
		return nil, err
	}
	return pi, run(ctx, pi)
}

func ContinuePIByKey(ctx context.Context, piKey string) (engine_types.ProcessInstance, error) {
	pi, err := getPI(piKey)
	if err != nil {
		return nil, err
	}
	return pi, run(ctx, pi)
}

// 基于消息名称对流程进行重启
func PublishEventForInstanceContinue(
	ctx context.Context,
	piKey string,
	messageName string,
	variables map[string]any,
) (engine_types.ProcessInstance, error) {
	pi, err := getPI(piKey)
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
			if _, err := pi.GetMessageSubscriptionManager().Create(
				ctx,
				msg,
			); err != nil {
				return nil, err
			}
			pi.SetVariables(variables)
			return pi, run(ctx, pi)
		}
	}
	return nil, fmt.Errorf("创建消息订阅失败")
}

// 创建新的流程实例
func CreatePIByByte(raw []byte, variables map[string]any) (engine_types.ProcessInstance, error) {
	if Engine().createHandler == nil {
		return nil, fmt.Errorf("未注册流程实例创建函数")
	}
	pi := Engine().createHandler(raw, variables)
	registerPI(pi)
	return pi, nil
}

// 通过流程ID创建实例
func CreatePIByID(processID string, variables map[string]any) (engine_types.ProcessInstance, error) {
	process, ok := Engine().processes[processID]
	if !ok {
		return nil, fmt.Errorf("流程ID【%s】未存在", processID)
	}
	return CreatePIByByte(process.raw, variables)
}
