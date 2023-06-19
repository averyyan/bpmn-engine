package memory_activity_manager

import (
	"context"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	"github.com/averyyan/bpmn-engine/memory/activity"
)

func New() *ActivityManager {
	return &ActivityManager{
		activs: make(map[string]*activity.Activity),
	}
}

// manager *ActivityManager github.com/averyyan/bpmn-engine/bpmn/engine/types.ActivityManager
type ActivityManager struct {
	activs map[string]*activity.Activity
}

// 创建新的活动
func (manager *ActivityManager) Create(ctx context.Context, pi engine_types.ProcessInstance, baseElement sepc_types.BaseElement) (engine_types.Activity, error) {
	activ := activity.New(pi, baseElement)
	manager.activs[activ.GetKey()] = activ
	return activ, nil
}

// 按状态查找活动
func (manager *ActivityManager) FindByStates(ctx context.Context, pi engine_types.ProcessInstance, states []sepc_element_types.LifecycleState) ([]engine_types.Activity, error) {
	var activs []engine_types.Activity
	for _, state := range states {
		for _, activ := range manager.activs {
			if state == activ.GetState() {
				activs = append(activs, activ)
			}
		}
	}
	return activs, nil
}

// 设置活动状态为完成
func (manager *ActivityManager) SetCompleted(ctx context.Context, activ engine_types.Activity) error {
	activited := manager.activs[activ.GetKey()]
	activited.SetState(sepc_element_types.Completed)
	return nil
}

// 设置活动状态为激活
func (manager *ActivityManager) SetActive(ctx context.Context, activ engine_types.Activity) error {
	activited := manager.activs[activ.GetKey()]
	activited.SetState(sepc_element_types.Active)
	return nil
}

// 设置活动状态为失败
func (manager *ActivityManager) SetFailed(ctx context.Context, activ engine_types.Activity, reason string) error {
	activited := manager.activs[activ.GetKey()]
	activited.SetState(sepc_element_types.Failed)
	return nil
}
