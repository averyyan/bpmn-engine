package engine

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/definitions"
	bpmn_util "github.com/averyyan/bpmn-engine/bpmn/util"
)

func newProcess(raw []byte) *process {
	return &process{
		raw: raw,
	}
}

// process *process github.com/averyyan/bpmn-engine/bpmn/engine/types.Process
type process struct {
	raw         []byte
	definitions *definitions.TDefinitions
}

// 获取流程数据
func (process *process) GetRaw() []byte {
	return process.raw
}

// 获取流程详情
func (process *process) GetDefinitions() (*definitions.TDefinitions, error) {
	if process.definitions == nil {
		definitions, err := bpmn_util.LoadFromByte(process.raw)
		if err != nil {
			return nil, err
		}
		process.definitions = definitions
	}
	return process.definitions, nil
}

// 获取流程ID
func (process *process) GetProcessID() (string, error) {
	definitions, err := process.GetDefinitions()
	if err != nil {
		return "", err
	}
	return definitions.Process.ID, nil
}
