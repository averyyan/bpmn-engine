package memory_process_instance_manager

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	memory_process_instance "github.com/averyyan/bpmn-engine/memory/process_instance"
)

func New() *ProcessInstanceManager {
	return &ProcessInstanceManager{
		ProcessInstances: make(map[string]*memory_process_instance.ProcessInstance),
	}
}

// manager *ProcessInstanceManager github.com/averyyan/bpmn-engine/bpmn/engine/types.ProcessInstanceManager
type ProcessInstanceManager struct {
	ProcessInstances map[string]*memory_process_instance.ProcessInstance `json:"process_instances"` // 内存流程实例Map
}

// 创建流程实例 ctx 上下文用于未来事务 raw 流程文件数据流 variables 实例上下文
func (manager *ProcessInstanceManager) Create(ctx context.Context, raw []byte, variables map[string]any) (engine_types.ProcessInstance, error) {
	pi := memory_process_instance.New(raw, variables)
	manager.ProcessInstances[pi.GetKey()] = pi.(*memory_process_instance.ProcessInstance)
	return pi, nil
}

// 通过Key找到流程实例
func (manager *ProcessInstanceManager) FindOneByKey(ctx context.Context, piKey string) (engine_types.ProcessInstance, error) {
	if pi, ok := manager.ProcessInstances[piKey]; ok {
		return pi, nil
	} else {
		return nil, fmt.Errorf("未存在Key【%s】的流程实例", piKey)
	}
}
