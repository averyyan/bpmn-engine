package memory_process

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/definitions"
	bpmn_util "github.com/averyyan/bpmn-engine/bpmn/util"
)

func New(raw []byte) *Process {
	return &Process{
		raw: raw,
	}
}

// p *Process github.com/averyyan/bpmn-engine/bpmn/engine/types.Process
type Process struct {
	raw         []byte // 流程文件数据
	definitions *definitions.TDefinitions
}

// 获取流程数据
func (p *Process) GetRaw() []byte {
	return p.raw
}

// 获取流程详情
func (p *Process) GetDefinitions() (*definitions.TDefinitions, error) {
	if p.definitions == nil {
		definitions, err := bpmn_util.LoadFromByte(p.raw)
		if err != nil {
			return nil, err
		}
		p.definitions = definitions
	}
	return p.definitions, nil
}

func (p *Process) GetProcessID() (string, error) {
	dfs, err := p.GetDefinitions()
	if err != nil {
		return "", nil
	}
	return dfs.Process.ID, nil
}
