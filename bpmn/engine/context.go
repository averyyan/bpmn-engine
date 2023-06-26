package engine

import (
	"embed"
	"fmt"
	"sync"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	bpmn_util "github.com/averyyan/bpmn-engine/bpmn/util"
)

// 全局服务储存
var engineCtx *engineContext
var once sync.Once

func Engine() *engineContext {
	once.Do(func() {
		engineCtx = &engineContext{
			processes:           make(map[string]*process),
			instances:           make(map[string]engine_types.ProcessInstance),
			serviceTaskHandlers: make(map[string]engine_types.TaskHandler),
		}
	})
	return engineCtx
}

type createHandler func(raw []byte, variables map[string]any) engine_types.ProcessInstance

// 主要用于服务储存
type engineContext struct {
	processes           map[string]*process
	instances           map[string]engine_types.ProcessInstance
	serviceTaskHandlers map[string]engine_types.TaskHandler
	createHandler       createHandler
}

// 注册
func RegisterPICreateHandler(handler createHandler) {
	Engine().createHandler = handler
}

// 注册 Servcie Task 元素处理函数
func RegisterServiceTaskHandler(taskType string, fun engine_types.TaskHandler) {
	Engine().serviceTaskHandlers[taskType] = fun
}

// 获取 Servcie Task 元素处理函数
func GetServiceTaskHandler(taskType string) (engine_types.TaskHandler, error) {
	handler, ok := Engine().serviceTaskHandlers[taskType]
	if !ok {
		return nil, fmt.Errorf("未存在元素【%s】绑定函数", taskType)
	}
	return handler, nil
}

// 获取流程
func GetProcess(processID string) engine_types.Process {
	return Engine().processes[processID]
}

// 注册流程实例进入
func registerPI(pi engine_types.ProcessInstance) {
	Engine().instances[pi.GetKey()] = pi
}

// 通过流程实例Key获取流程实例
func getPI(piKey string) (engine_types.ProcessInstance, error) {
	pi, ok := Engine().instances[piKey]
	if !ok {
		return nil, fmt.Errorf("Key【%s】流程实例为存在", piKey)
	}
	return pi, nil
}

// 通过嵌入资源
func LoadFromEmbed(myfs *embed.FS) error {
	datas, err := bpmn_util.LoadFromEmbed(myfs)
	if err != nil {
		return err
	}
	for _, data := range datas {
		process := newProcess(data)
		processID, err := process.GetProcessID()
		if err != nil {
			return err
		}
		Engine().processes[processID] = process
	}
	return nil
}
