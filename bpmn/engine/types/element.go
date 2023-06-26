package engine_types

import (
	"context"

	"github.com/averyyan/bpmn-engine/bpmn/sepc/activity"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/event"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 未完成元素管理
type ElementManager interface {
	// 找到所有激活的阻塞元素
	FindActiveElements(ctx context.Context) []sepc_types.BaseElement
	// 创建阻塞消息中间事件
	CreateMessageICE(ctx context.Context, baseElement *event.TIntermediateCatchEvent) (BaseElement, error)
	// 删除阻塞消息中间件事件
	DeleteMessageICE(ctx context.Context, msgICEKey string) error
	// 创建阻塞消息中间事件
	CreateCallActivity(ctx context.Context, baseElement *activity.TCallActivity) (BaseElement, error)
	// 删除阻塞消息中间件事件
	DeleteCallActivity(ctx context.Context, callActivityKey string) error
	// 通过元素ID找到储存阻塞的消息中间事件
	FindOneMessageICE(ctx context.Context, elementID string) (BaseElement, error)
	// 通过元素ID找到储存阻塞的重复活动
	FindOneCallActivity(ctx context.Context, elementID string) (BaseElement, error)
}

type BaseElement interface {
	// 获取元素
	GetElement() sepc_types.BaseElement
	// 获取元素类型
	GetElementType() sepc_element_types.ElementType
	// 元素唯一Key
	GetKey() string
}
