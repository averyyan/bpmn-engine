package memory_ice_manager

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
	"github.com/averyyan/bpmn-engine/memory/ice"
)

func New() *ICEManager {
	return &ICEManager{
		ices: make(map[string]*ice.IntermediateCatchEvent),
	}
}

// manager *ICEManager github.com/averyyan/bpmn-engine/bpmn/engine/types.ICEManager
type ICEManager struct {
	ices map[string]*ice.IntermediateCatchEvent
}

// 创建消息中间事件
func (manager *ICEManager) CreateMsgICE(ctx context.Context, pi engine_types.ProcessInstance, baseElement sepc_types.IntermediateCatchEvent) (engine_types.IntermediateCatchEvent, error) {
	ice := ice.New(pi, baseElement)
	manager.ices[ice.GetKey()] = ice
	return ice, nil
}

// 按状态查找消息中间事件
func (manager *ICEManager) FindMsgICEByStates(ctx context.Context, pi engine_types.ProcessInstance, states []sepc_element_types.LifecycleState) ([]engine_types.IntermediateCatchEvent, error) {
	var ices []engine_types.IntermediateCatchEvent
	for _, ice := range manager.ices {
		if ice.GetProcessInstanceKey() == pi.GetKey() {
			for _, state := range states {
				if ice.GetState() == state {
					ices = append(ices, ice)
				}
			}
		}
	}
	return ices, nil
}

// 按照状态和ID找到消息中间事件
func (manager *ICEManager) FindOneMsgICEByStateAndID(ctx context.Context, pi engine_types.ProcessInstance, state sepc_element_types.LifecycleState, elementID string) (engine_types.IntermediateCatchEvent, error) {
	for _, ice := range manager.ices {
		if ice.GetProcessInstanceKey() == pi.GetKey() {
			if ice.GetState() == state && ice.GetElementID() == elementID {
				return ice, nil
			}
		}
	}
	return nil, fmt.Errorf("未找到ID【%s】的元素", elementID)
}

// 设置活动状态为完成
func (manager *ICEManager) SetCompleted(ctx context.Context, ice engine_types.IntermediateCatchEvent) error {
	_ice := manager.ices[ice.GetKey()]
	_ice.SetState(sepc_element_types.Completed)
	return nil
}

// 设置活动状态为失败
func (manager *ICEManager) SetFailed(ctx context.Context, ice engine_types.IntermediateCatchEvent) error {
	_ice := manager.ices[ice.GetKey()]
	_ice.SetState(sepc_element_types.Failed)
	return nil
}
