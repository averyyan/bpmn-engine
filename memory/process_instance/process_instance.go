package memory_process_instance

import (
	"fmt"

	"github.com/averyyan/bpmn-engine/bpmn/sepc/definitions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
	bpmn_util "github.com/averyyan/bpmn-engine/bpmn/util"
)

func New(key string, raw []byte, variables map[string]any) *ProcessInstance {
	return &ProcessInstance{
		state:          sepc_pi_types.Ready,
		key:            key,
		raw:            raw,
		variables:      variables,
		ScheduledFlows: make([]string, 0),
	}
}

// pi *ProcessInstance github.com/averyyan/bpmn-engine/bpmn/engine/types.ProcessInstance
type ProcessInstance struct {
	// 流程实例状态
	state sepc_pi_types.State
	// 流程实例唯一Key
	key string
	// 流程内容
	raw []byte
	// 流程实例上下文
	variables map[string]any
	// 流程实例内容
	definitions *definitions.TDefinitions
	// 流程序列流
	ScheduledFlows []string
	// 元素Map
	elementMap map[string]sepc_types.BaseElement
}

// 获取流程唯一Key
func (pi *ProcessInstance) GetKey() string {
	return pi.key
}

// 获取流程实例状态
func (pi *ProcessInstance) GetState() sepc_pi_types.State {
	return pi.state
}

func (pi *ProcessInstance) SetState(state sepc_pi_types.State) {
	pi.state = state
}

// 获取流程详情
func (pi *ProcessInstance) GetDefinitions() (*definitions.TDefinitions, error) {
	if pi.definitions == nil {
		definitions, err := bpmn_util.LoadFromByte(pi.raw)
		if err != nil {
			return nil, err
		}
		pi.definitions = definitions
	}
	return pi.definitions, nil
}

// 通过元素ID获取流程实例中元素
func (pi *ProcessInstance) FindBaseElementsById(id string) (sepc_types.BaseElement, error) {
	if pi.elementMap == nil {
		definitions, err := pi.GetDefinitions()
		if err != nil {
			return nil, err
		}
		pi.elementMap = bpmn_util.MakeElementMap(definitions)
	}
	if pi.elementMap[id] == nil {
		return nil, fmt.Errorf("元素【%s】未存在", id)
	}
	return pi.elementMap[id], nil
}
