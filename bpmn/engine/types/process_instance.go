package engine_types

import (
	"context"

	"github.com/averyyan/bpmn-engine/bpmn/sepc/definitions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
)

// 流程实例管理接口
type ProcessInstanceManager interface {
	// 创建流程实例 ctx 上下文用于未来事务 raw 流程文件数据流 variables 实例上下文
	Create(ctx context.Context, raw []byte, variables map[string]any) (ProcessInstance, error)
	// 设置流程实例为激活状态
	SetActive(ctx context.Context, pi ProcessInstance) error
	// 设置流程实例为完成状态
	SetCompleted(ctx context.Context, pi ProcessInstance) error
	// 设置流程实例为失败状态
	SetFailed(ctx context.Context, pi ProcessInstance) error
	// 添加flowID到ScheduledFlows
	AppendScheduledFlows(ctx context.Context, pi ProcessInstance, flowID string) error
	// 删除flowID到ScheduledFlows
	RemoveScheduledFlows(ctx context.Context, pi ProcessInstance, flowID string) error
}

// 流程实例接口
type ProcessInstance interface {
	// 获取流程唯一Key
	GetKey() string
	// 获取流程实例状态
	GetState() sepc_pi_types.State
	// 获取流程详情
	GetDefinitions() (*definitions.TDefinitions, error)
	// 通过元素ID获取流程实例中元素
	FindBaseElementsById(id string) (sepc_types.BaseElement, error)
}
