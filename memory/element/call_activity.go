package memory_element

import (
	"context"

	"github.com/averyyan/bpmn-engine/bpmn/sepc/activity"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

func NewCallActivity(element *activity.TCallActivity) *CallActivity {
	return &CallActivity{
		Element:     element,
		BaseElement: New(element),
	}
}

// element *CallActivity github.com/averyyan/bpmn-engine/bpmn/engine/types.CallActivity
type CallActivity struct {
	BaseElement             `json:",inline"`
	Element                 *activity.TCallActivity `json:"element"`
	ChildProcessInstanceKey string                  `json:"child_process_instance_key"`
}

// 获取元素
func (element *CallActivity) GetElement() sepc_types.BaseElement {
	return element.Element
}

// 设置子流程
func (element *CallActivity) SetChildPIKey(ctx context.Context, piKey string) error {
	element.ChildProcessInstanceKey = piKey
	return nil
}

// 获取子流程
func (element *CallActivity) GetChildPIKey() string {
	return element.ChildProcessInstanceKey
}
