package memory_process_instance

import (
	"context"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/definitions"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
	bpmn_util "github.com/averyyan/bpmn-engine/bpmn/util"
	memory_element_manager "github.com/averyyan/bpmn-engine/memory/element/manager"
	memory_msgsub_manager "github.com/averyyan/bpmn-engine/memory/msgsub/manager"
	"github.com/segmentio/ksuid"
)

func New(raw []byte, variables map[string]any) engine_types.ProcessInstance {
	return &ProcessInstance{
		Raw:                        raw,
		State:                      sepc_pi_types.Ready,
		Key:                        ksuid.New().String(),
		Variables:                  variables,
		ScheduledFlows:             make([]string, 0),
		ElementManager:             memory_element_manager.New(),
		MessageSubscriptionManager: memory_msgsub_manager.New(),
	}
}

// pi *ProcessInstance github.com/averyyan/bpmn-engine/bpmn/engine/types.ProcessInstance
type ProcessInstance struct {
	Raw                        []byte                                            `json:"raw"`                         // 流程内容
	State                      sepc_pi_types.State                               `json:"state"`                       // 流程实例状态
	Key                        string                                            `json:"key"`                         // 流程实例唯一Key
	Variables                  map[string]any                                    `json:"variables"`                   // 流程实例上下文
	ScheduledFlows             []string                                          `json:"scheduled_flows"`             // 流程序列流
	ElementManager             *memory_element_manager.ElementManager            `json:"element_manager"`             // 阻塞元素管理
	MessageSubscriptionManager *memory_msgsub_manager.MessageSubscriptionManager `json:"msgsub_manager"`              // 消息订阅
	ParentProcessInstanceKey   string                                            `json:"parent_process_instance_key"` // 父级流程实例Key
	definitions                *definitions.TDefinitions                         // 流程实例内容
}

// 设置上下文参数
func (pi *ProcessInstance) SetVariables(variables map[string]any) error {
	if pi.Variables == nil {
		pi.Variables = make(map[string]any)
	}
	for key, value := range variables {
		pi.Variables[key] = value
	}
	return nil
}

// 获取阻塞元素管理
func (pi *ProcessInstance) GetElementManager() engine_types.ElementManager {
	return pi.ElementManager
}

// 获取消息订阅管理
func (pi *ProcessInstance) GetMessageSubscriptionManager() engine_types.MessageSubscriptionManager {
	return pi.MessageSubscriptionManager
}

// 获取流程唯一Key
func (pi *ProcessInstance) GetKey() string {
	return pi.Key
}

// 获取流程实例状态
func (pi *ProcessInstance) GetState() sepc_pi_types.State {
	return pi.State
}

// 获取流程详情
func (pi *ProcessInstance) GetDefinitions() (*definitions.TDefinitions, error) {
	if pi.definitions == nil {
		definitions, err := bpmn_util.LoadFromByte(pi.Raw)
		if err != nil {
			return nil, err
		}
		pi.definitions = definitions
	}
	return pi.definitions, nil
}

// 获取流程实例全局上下文
func (pi *ProcessInstance) GetVariables() map[string]any {
	return pi.Variables
}

// 获取父级流程实例Key
func (pi *ProcessInstance) GetParentProcessInstanceKey() string {
	return pi.ParentProcessInstanceKey
}

// 设置父级流程实例
func (pi *ProcessInstance) SetParentProcessInstanceKey(ctx context.Context, key string) error {
	pi.ParentProcessInstanceKey = key
	return nil
}

// 设置流程实例为激活状态
func (pi *ProcessInstance) SetActive(ctx context.Context) error {
	pi.State = sepc_pi_types.Active
	return nil
}

// 设置流程实例为完成状态
func (pi *ProcessInstance) SetCompleted(ctx context.Context) error {
	pi.State = sepc_pi_types.Completed
	return nil
}

// 设置流程实例为失败状态
func (pi *ProcessInstance) SetFailed(ctx context.Context) error {
	pi.State = sepc_pi_types.Failed
	return nil
}

// 判断是否存在
func (pi *ProcessInstance) HasScheduledFlow(ctx context.Context, flowID string) bool {
	return bpmn_util.ContainsString(pi.ScheduledFlows, flowID)
}

// 获取
func (pi *ProcessInstance) GetScheduledFlows(ctx context.Context) []string {
	return pi.ScheduledFlows
}

// 添加flowID到ScheduledFlows 处理并行网关使用
func (pi *ProcessInstance) AppendScheduledFlow(ctx context.Context, flowID string) error {
	pi.ScheduledFlows = append(pi.ScheduledFlows, flowID)
	return nil
}

// 删除flowID到ScheduledFlows 处理并行网关使用
func (pi *ProcessInstance) RemoveScheduledFlow(ctx context.Context, flowID string) error {
	pi.ScheduledFlows = bpmn_util.RemoveString(pi.ScheduledFlows, flowID)
	return nil
}
