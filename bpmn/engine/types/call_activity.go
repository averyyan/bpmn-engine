package engine_types

import "context"

type CallActivityManager interface {
	SetChildPIKey(ctx context.Context, base BaseElement, piKey string) error
}

type CallActivity interface {
	GetChildProcessInstanceKey() string // 获取子流程
}
