package memory_call_activity_manager

import (
	"context"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	"github.com/averyyan/bpmn-engine/memory/call_activity"
	memory_manager "github.com/averyyan/bpmn-engine/memory/manager"
)

func New() *CallActivityManager {
	return &CallActivityManager{
		BaseManager: memory_manager.New[call_activity.CallActivity](),
	}
}

// manager *CallActivityManager github.com/averyyan/bpmn-engine/bpmn/engine/types.BaseManager
// manager *CallActivityManager github.com/averyyan/bpmn-engine/bpmn/engine/types.CallActivityManager
type CallActivityManager struct {
	*memory_manager.BaseManager[call_activity.CallActivity]
}

// 创建新的活动
func (manager *CallActivityManager) Create(ctx context.Context, pi engine_types.ProcessInstance, baseElement sepc_types.BaseElement) (engine_types.BaseElement, error) {
	data := call_activity.New(pi, baseElement)
	manager.Add(data, data)
	return data, nil
}

func (manager *CallActivityManager) SetChildPIKey(ctx context.Context, base engine_types.BaseElement, piKey string) error {
	data := manager.GetValue(base.GetKey())
	data.SetChildProcessInstanceKey(piKey)
	return nil
}

// 设置活动状态为完成
func (manager *CallActivityManager) SetCompleted(ctx context.Context, base engine_types.BaseElement) error {
	data := manager.GetValue(base.GetKey())
	data.SetState(sepc_element_types.Completed)
	return nil
}

// 设置活动状态为激活
func (manager *CallActivityManager) SetActive(ctx context.Context, base engine_types.BaseElement) error {
	data := manager.GetValue(base.GetKey())
	data.SetState(sepc_element_types.Active)
	return nil
}

// 设置活动状态为失败
func (manager *CallActivityManager) SetFailed(ctx context.Context, base engine_types.BaseElement, reason string) error {
	data := manager.GetValue(base.GetKey())
	data.SetState(sepc_element_types.Failed)
	return nil
}
