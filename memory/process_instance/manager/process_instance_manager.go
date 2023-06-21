package memory_process_instance_manager

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
	bpmn_util "github.com/averyyan/bpmn-engine/bpmn/util"
	memory_process_instance "github.com/averyyan/bpmn-engine/memory/process_instance"
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

// 通过Key找到流程实例
func (manager *ProcessInstanceManager) FindOneByKey(ctx context.Context, piKey string) (engine_types.ProcessInstance, error) {
	if pi, ok := manager.pis[piKey]; ok {
		return pi, nil
	} else {
		return nil, fmt.Errorf("未存在Key【%s】的流程实例", piKey)
	}
}

// 创建流程实例 ctx 上下文用于未来事务 raw 流程文件数据流 variables 实例上下文
func (manager *ProcessInstanceManager) Create(ctx context.Context, raw []byte, variables map[string]any) (engine_types.ProcessInstance, error) {
	pi := memory_process_instance.New(raw, variables)
	manager.pis[pi.GetKey()] = pi
	return pi, nil
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
	_pi.ScheduledFlows = append(_pi.ScheduledFlows, flowID)
	return nil
}

// 删除flowID到ScheduledFlows
func (manager *ProcessInstanceManager) RemoveScheduledFlows(ctx context.Context, pi engine_types.ProcessInstance, flowID string) error {
	_pi := manager.pis[pi.GetKey()]
	_pi.ScheduledFlows = bpmn_util.RemoveString(_pi.ScheduledFlows, flowID)
	return nil
}

// 判断是否存在
func (manager *ProcessInstanceManager) HasScheduledFlow(ctx context.Context, pi engine_types.ProcessInstance, flowID string) bool {
	_pi := manager.pis[pi.GetKey()]
	return bpmn_util.ContainsString(_pi.ScheduledFlows, flowID)
}

// 获取
func (manager *ProcessInstanceManager) GetScheduledFlows(ctx context.Context, pi engine_types.ProcessInstance) []string {
	_pi := manager.pis[pi.GetKey()]
	return _pi.ScheduledFlows
}
