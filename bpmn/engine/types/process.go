package engine_types

import (
	"context"
	"embed"

	"github.com/averyyan/bpmn-engine/bpmn/sepc/definitions"
)

// 流程管理接口
type ProcessManager interface {
	// 创建新流程
	Create(ctx context.Context, raw []byte) (Process, error)
	// 通过流程ID找到对应流程
	FindOneByID(ctx context.Context, processID string) (Process, error)
	// 读取文件夹
	LoadFromEmbed(ctx context.Context, myfs *embed.FS) error
}

// 流程接口
type Process interface {
	// 获取流程数据
	GetRaw() []byte
	// 获取流程详情
	GetDefinitions() (*definitions.TDefinitions, error)
	// 获取流程ID
	GetProcessID() (string, error)
}
