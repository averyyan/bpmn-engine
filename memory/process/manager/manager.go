package memory_process_manager

import (
	"context"
	"fmt"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
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

func (manager *ProcessManager) FindOneByID(ctx context.Context, processID string) (engine_types.Process, error) {
	if p, ok := manager.ps[processID]; ok {
		return p, nil
	} else {
		return nil, fmt.Errorf("未存在ID【%s】的流程", processID)
	}
}
