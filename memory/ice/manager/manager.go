package memory_ice_manager

import (
	"context"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	"github.com/averyyan/bpmn-engine/memory/ice"
	memory_manager "github.com/averyyan/bpmn-engine/memory/manager"
)

func New() *ICEManager {
	return &ICEManager{
		BaseManager: memory_manager.New[ice.IntermediateCatchEvent](),
	}
}

// manager *ICEManager github.com/averyyan/bpmn-engine/bpmn/engine/types.BaseManager
// manager *ICEManager github.com/averyyan/bpmn-engine/bpmn/engine/types.ICEManager
type ICEManager struct {
	*memory_manager.BaseManager[ice.IntermediateCatchEvent]
}

// 创建新的活动
func (manager *ICEManager) Create(ctx context.Context, pi engine_types.ProcessInstance, baseElement sepc_types.BaseElement) (engine_types.BaseElement, error) {
	ice := ice.New(pi, baseElement)
	manager.Add(ice, ice)
	return ice, nil
}

// 设置活动状态为完成
func (manager *ICEManager) SetCompleted(ctx context.Context, base engine_types.BaseElement) error {
	data := manager.GetValue(base.GetKey())
	data.SetState(sepc_element_types.Completed)
	return nil
}

// 设置活动状态为激活
func (manager *ICEManager) SetActive(ctx context.Context, base engine_types.BaseElement) error {
	data := manager.GetValue(base.GetKey())
	data.SetState(sepc_element_types.Active)
	return nil
}

// 设置活动状态为失败
func (manager *ICEManager) SetFailed(ctx context.Context, base engine_types.BaseElement, reason string) error {
	data := manager.GetValue(base.GetKey())
	data.SetState(sepc_element_types.Failed)
	return nil
}
