package engine

import (
	"encoding/json"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
)

// 保存为json文件
func SaveProcessInstance(pi engine_types.ProcessInstance) ([]byte, error) {
	return json.Marshal(pi)
}

// 导入json文件
func LoadProcessInstance[T engine_types.ProcessInstance](raw []byte) (engine_types.ProcessInstance, error) {
	var pi *T
	if err := json.Unmarshal(raw, &pi); err != nil {
		return nil, err
	}
	registerPI(*pi)
	return *pi, nil
}
