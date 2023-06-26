package memory_process_manager

import (
	"context"
	"embed"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	bpmn_util "github.com/averyyan/bpmn-engine/bpmn/util"
	memory_process "github.com/averyyan/bpmn-engine/memory/process"
)

func New() *ProcessManager {
	return &ProcessManager{
		ps: make(map[string]*memory_process.Process),
	}
}

// manager *ProcessManager github.com/averyyan/bpmn-engine/bpmn/engine/types.ProcessManager
type ProcessManager struct {
	ps map[string]*memory_process.Process // 内存流程Map
}

func (manager *ProcessManager) Create(ctx context.Context, raw []byte) (engine_types.Process, error) {
	p := memory_process.New(raw)
	processID, err := p.GetProcessID()
	if err != nil {
		return nil, err
	}
	manager.ps[processID] = p
	return p, nil
}

// 读取文件夹
func (manager *ProcessManager) LoadFromEmbed(ctx context.Context, myfs *embed.FS) error {
	processes, err := bpmn_util.LoadFromEmbed(myfs)
	if err != nil {
		return err
	}
	for _, process := range processes {
		if _, err := manager.Create(ctx, process); err != nil {
			return err
		}
	}
	return nil
}

func (manager *ProcessManager) FindOneByID(ctx context.Context, processID string) (engine_types.Process, error) {
	if p, ok := manager.ps[processID]; ok {
		return p, nil
	} else {
		return nil, fmt.Errorf("未存在ID【%s】的流程", processID)
	}
}
