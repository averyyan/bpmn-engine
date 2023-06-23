package memory_activity_manager

import (
	"context"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	"github.com/averyyan/bpmn-engine/memory/activity"
	memory_manager "github.com/averyyan/bpmn-engine/memory/manager"
)

func New() *ActivityManager {
	return &ActivityManager{
		BaseManager: memory_manager.New[activity.Activity](),
	}
}

// manager *ActivityManager github.com/averyyan/bpmn-engine/bpmn/engine/types.BaseManager
// manager *ActivityManager github.com/averyyan/bpmn-engine/bpmn/engine/types.ActivityManager
type ActivityManager struct {
	*memory_manager.BaseManager[activity.Activity]
}

// 创建新的活动
func (manager *ActivityManager) Create(ctx context.Context, pi engine_types.ProcessInstance, baseElement sepc_types.BaseElement) (engine_types.BaseElement, error) {
	activ := activity.New(pi, baseElement)
	manager.Add(activ, activ)
	return activ, nil
}

// 按状态查找活动
// func (manager *ActivityManager) FindByStates(ctx context.Context, pi engine_types.ProcessInstance, states []sepc_element_types.LifecycleState) ([]engine_types.BaseElement, error) {
// 	return memory_manager.NewSearch(*manager.BaseManager).FilterPI(pi).FilterStates(states).GetDatas(), nil
// }

// 按元素ID和元素状态查找元素
// func (manager *ActivityManager) FindOneByStateAndID(ctx context.Context, pi engine_types.ProcessInstance, state sepc_element_types.LifecycleState, elementID string) (engine_types.BaseElement, error) {
// 	element := memory_manager.NewSearch(*manager.BaseManager).FilterPI(pi).FilterStates(
// 		[]sepc_element_types.LifecycleState{sepc_element_types.Active},
// 	).FindOneByID(elementID)
// 	if element == nil {
// 		return nil, fmt.Errorf("未找到ID【%s】元素", elementID)
// 	}
// 	return element, nil
// }

// 设置活动状态为完成
func (manager *ActivityManager) SetCompleted(ctx context.Context, base engine_types.BaseElement) error {
	activited := manager.GetValue(base.GetKey())
	activited.SetState(sepc_element_types.Completed)
	return nil
}

// 设置活动状态为激活
func (manager *ActivityManager) SetActive(ctx context.Context, base engine_types.BaseElement) error {
	activited := manager.GetValue(base.GetKey())
	activited.SetState(sepc_element_types.Active)
	return nil
}

// 设置活动状态为失败
func (manager *ActivityManager) SetFailed(ctx context.Context, base engine_types.BaseElement, reason string) error {
	activited := manager.GetValue(base.GetKey())
	activited.SetState(sepc_element_types.Failed)
	return nil
}
