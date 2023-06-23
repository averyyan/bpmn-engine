package memory_manager

import (
	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

func NewSearch[T any](manager BaseManager[T]) *searchManager[T] {
	return &searchManager[T]{
		BaseManager: &manager,
	}
}

type searchManager[T any] struct {
	*BaseManager[T]
}

func (s *searchManager[T]) GetDatas() []engine_types.BaseElement {
	return s.datas
}

func (s *searchManager[T]) FilterPI(pi engine_types.ProcessInstance) *searchManager[T] {
	var _datas []engine_types.BaseElement
	for _, data := range s.datas {
		if data.GetProcessInstanceKey() == pi.GetKey() {
			_datas = append(_datas, data)
		}
	}
	s.datas = _datas
	return s
}

func (s *searchManager[T]) FilterStates(states []sepc_element_types.LifecycleState) *searchManager[T] {
	var _datas []engine_types.BaseElement
	for _, data := range s.datas {
		for _, state := range states {
			if state == data.GetState() {
				_datas = append(_datas, data)
			}
		}
	}
	s.datas = _datas
	return s
}

func (s *searchManager[T]) FindOneByID(elementID string) engine_types.BaseElement {
	for _, data := range s.datas {
		if data.GetElementID() == elementID {
			return data
		}
	}
	return nil
}
