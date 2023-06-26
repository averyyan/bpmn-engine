package engine_types

import "context"

type CallActivityManager interface {
}

type CallActivity interface {
	// 设置子流程
	SetChildPIKey(ctx context.Context, piKey string) error
	// 获取子流程
	GetChildPIKey() string
}
