package memory_element_manager

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/activity"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/event"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	memory_element "github.com/averyyan/bpmn-engine/memory/element"
)

func New() *ElementManager {
	return &ElementManager{
		MsgICEElements:       make(map[string]*memory_element.MessageICE),
		CallActivityElements: make(map[string]*memory_element.CallActivity),
	}
}

// 阻塞元素管理
// manager *ElementManager github.com/averyyan/bpmn-engine/bpmn/engine/types.ElementManager
type ElementManager struct {
	MsgICEElements       map[string]*memory_element.MessageICE   `json:"msg_ice_elements"`
	CallActivityElements map[string]*memory_element.CallActivity `json:"call_activity_elements"`
}

// 创建阻塞消息中间事件
func (manager *ElementManager) CreateCallActivity(ctx context.Context, baseElement *activity.TCallActivity) (engine_types.BaseElement, error) {
	cacitv := memory_element.NewCallActivity(baseElement)
	manager.CallActivityElements[cacitv.GetKey()] = cacitv
	return cacitv, nil
}

// 删除阻塞消息中间件事件
func (manager *ElementManager) DeleteCallActivity(ctx context.Context, callActivityKey string) error {
	delete(manager.CallActivityElements, callActivityKey)
	return nil
}

// 创建阻塞消息中间事件
func (manager *ElementManager) CreateMessageICE(ctx context.Context, baseElement *event.TIntermediateCatchEvent) (engine_types.BaseElement, error) {
	ice := memory_element.NewMessageICE(baseElement)
	manager.MsgICEElements[ice.GetKey()] = ice
	return ice, nil
}

// 删除阻塞消息中间件事件
func (manager *ElementManager) DeleteMessageICE(ctx context.Context, msgICEKey string) error {
	delete(manager.MsgICEElements, msgICEKey)
	return nil
}

// 找到所有激活的阻塞元素
func (manager *ElementManager) FindActiveElements(ctx context.Context) []sepc_types.BaseElement {
	var baseElements []sepc_types.BaseElement
	for _, element := range manager.MsgICEElements {
		baseElements = append(baseElements, element.GetElement())
	}
	for _, element := range manager.CallActivityElements {
		baseElements = append(baseElements, element.GetElement())
	}
	return baseElements
}

// 通过元素ID找到储存阻塞的消息中间事件
func (manager *ElementManager) FindOneMessageICE(ctx context.Context, elementID string) (engine_types.BaseElement, error) {
	var elements []engine_types.BaseElement
	for _, element := range manager.MsgICEElements {
		elements = append(elements, element)
	}
	return manager.findActivedOneByID(ctx, elements, elementID)
}

// 通过元素ID找到储存阻塞的重复活动
func (manager *ElementManager) FindOneCallActivity(ctx context.Context, elementID string) (engine_types.BaseElement, error) {
	var elements []engine_types.BaseElement
	for _, element := range manager.CallActivityElements {
		elements = append(elements, element)
	}
	return manager.findActivedOneByID(ctx, elements, elementID)
}

func (manager *ElementManager) findActivedOneByID(ctx context.Context, elements []engine_types.BaseElement, elementID string) (engine_types.BaseElement, error) {
	for _, element := range elements {
		if element.GetElement().GetID() == elementID {
			return element, nil
		}
	}
	return nil, fmt.Errorf("未找到元素ID【%s】的阻塞元素", elementID)
}
