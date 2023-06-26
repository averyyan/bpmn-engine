package engine_types

import (
	"context"

	"github.com/averyyan/bpmn-engine/bpmn/sepc/definitions"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
)

// 流程实例管理接口
type ProcessInstanceManager interface {
	// 创建流程实例 ctx 上下文用于未来事务 raw 流程文件数据流 variables 实例上下文
	Create(ctx context.Context, raw []byte, variables map[string]any) (ProcessInstance, error)
	// 通过Key找到流程实例
	FindOneByKey(ctx context.Context, piKey string) (ProcessInstance, error)
}

// 流程实例接口
type ProcessInstance interface {
	// 获取流程唯一Key
	GetKey() string
	// 获取流程实例状态
	GetState() sepc_pi_types.State
	// 获取流程详情
	GetDefinitions() (*definitions.TDefinitions, error)
	// 获取流程实例全局上下文
	GetVariables() map[string]any
	// 获取父级流程实例Key
	GetParentProcessInstanceKey() string
	// 获取阻塞元素管理
	GetElementManager() ElementManager
	// 获取消息订阅管理
	GetMessageSubscriptionManager() MessageSubscriptionManager
	// 设置上下文参数
	SetVariables(variables map[string]any) error
	// 设置父级流程实例
	SetParentProcessInstanceKey(ctx context.Context, key string) error
	// 设置流程实例为激活状态
	SetActive(ctx context.Context) error
	// 设置流程实例为完成状态
	SetCompleted(ctx context.Context) error
	// 设置流程实例为失败状态
	SetFailed(ctx context.Context) error
	// 判断是否存在
	HasScheduledFlow(ctx context.Context, flowID string) bool
	// 获取
	GetScheduledFlows(ctx context.Context) []string
	// 添加flowID到ScheduledFlows 处理并行网关使用
	AppendScheduledFlow(ctx context.Context, flowID string) error
	// 删除flowID到ScheduledFlows 处理并行网关使用
	RemoveScheduledFlow(ctx context.Context, flowID string) error
}
