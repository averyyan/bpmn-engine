package memory_process_instance

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/definitions"
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
}

// 获取流程实例全局上下文
func (pi *ProcessInstance) GetVariables() map[string]any {
	return pi.variables
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
