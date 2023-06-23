package memory_manager

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

func New[T any]() *BaseManager[T] {
	return &BaseManager[T]{
		datas:   make([]engine_types.BaseElement, 0),
		dataMap: make(map[string]*T),
	}
}

// manager *BaseManager github.com/averyyan/bpmn-engine/bpmn/engine/types.BaseManager
type BaseManager[T any] struct {
	datas   []engine_types.BaseElement
	dataMap map[string]*T
}

// 按状态查找活动
func (manager *BaseManager[T]) FindByStates(ctx context.Context, pi engine_types.ProcessInstance, states []sepc_element_types.LifecycleState) ([]engine_types.BaseElement, error) {
	return NewSearch(*manager).FilterPI(pi).FilterStates(states).GetDatas(), nil
}

// 按元素ID和元素状态查找元素
func (manager *BaseManager[T]) FindOneByStateAndID(ctx context.Context, pi engine_types.ProcessInstance, state sepc_element_types.LifecycleState, elementID string) (engine_types.BaseElement, error) {
	element := NewSearch(*manager).FilterPI(pi).FilterStates(
		[]sepc_element_types.LifecycleState{state},
	).FindOneByID(elementID)
	if element == nil {
		return nil, fmt.Errorf("未找到ID【%s】元素", elementID)
	}
	return element, nil
}

func (manager *BaseManager[T]) Add(data engine_types.BaseElement, value *T) {
	if _, ok := manager.dataMap[data.GetKey()]; ok {
		for i, _data := range manager.datas {
			if _data.GetKey() == data.GetKey() {
				manager.datas[i] = data
			}
		}
	} else {
		manager.datas = append(manager.datas, data)
	}
	manager.dataMap[data.GetKey()] = value
}

func (manager *BaseManager[T]) Delete(data engine_types.BaseElement) {
	manager.datas = deleteItem(manager.datas, data)
	delete(manager.dataMap, data.GetKey())
}

func (manager *BaseManager[T]) GetValue(key string) *T {
	return manager.dataMap[key]
}

func (manager *BaseManager[T]) GetData(key string) engine_types.BaseElement {
	for _, data := range manager.datas {
		if data.GetKey() == key {
			return data
		}
	}
	return nil
}

func deleteItem(items []engine_types.BaseElement, value engine_types.BaseElement) []engine_types.BaseElement {
	i := 0
	for _, item := range items {
		if item.GetKey() != value.GetKey() {
			items[i] = item
			i++
		}
	}
	return items[:i]
}
