package memory_process_instance_manager

import (
	"context"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
	bpmn_util "github.com/averyyan/bpmn-engine/bpmn/util"
	memory_process_instance "github.com/averyyan/bpmn-engine/memory/process_instance"
	"github.com/segmentio/ksuid"
)

func New() *ProcessInstanceManager {
	return &ProcessInstanceManager{
		pis: make(map[string]*memory_process_instance.ProcessInstance),
	}
}

// manager *ProcessInstanceManager github.com/averyyan/bpmn-engine/bpmn/engine/types.ProcessInstanceManager
type ProcessInstanceManager struct {
	pis map[string]*memory_process_instance.ProcessInstance // 内存流程实例Map
}

// 创建流程实例 ctx 上下文用于未来事务 raw 流程文件数据流 variables 实例上下文
func (manager *ProcessInstanceManager) Create(ctx context.Context, raw []byte, variables map[string]any) (engine_types.ProcessInstance, error) {
	key := ksuid.New().String()
	manager.pis[key] = memory_process_instance.New(key, raw, variables)
	return manager.pis[key], nil
}

// 设置流程实例为激活状态
func (manager *ProcessInstanceManager) SetActive(ctx context.Context, pi engine_types.ProcessInstance) error {
	manager.pis[pi.GetKey()].SetState(sepc_pi_types.Active)
	return nil
}

// 设置流程实例为完成状态
func (manager *ProcessInstanceManager) SetCompleted(ctx context.Context, pi engine_types.ProcessInstance) error {
	manager.pis[pi.GetKey()].SetState(sepc_pi_types.Completed)
	return nil
}

// 设置流程实例为失败状态
func (manager *ProcessInstanceManager) SetFailed(ctx context.Context, pi engine_types.ProcessInstance) error {
	manager.pis[pi.GetKey()].SetState(sepc_pi_types.Failed)
	return nil
}

// 添加flowID到ScheduledFlows
func (manager *ProcessInstanceManager) AppendScheduledFlows(ctx context.Context, pi engine_types.ProcessInstance, flowID string) error {
	_pi := manager.pis[pi.GetKey()]
	_pi.ScheduledFlows = bpmn_util.RemoveString(_pi.ScheduledFlows, flowID)
	return nil
}

// 删除flowID到ScheduledFlows
func (manager *ProcessInstanceManager) RemoveScheduledFlows(ctx context.Context, pi engine_types.ProcessInstance, flowID string) error {
	_pi := manager.pis[pi.GetKey()]
	_pi.ScheduledFlows = bpmn_util.RemoveString(_pi.ScheduledFlows, flowID)
	return nil
}
