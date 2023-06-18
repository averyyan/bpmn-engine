package engine_types

import (
	"time"

	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

type BaseElement interface {
	GetElementID() string
	GetElementName() string
	GetKey() string
	GetProcessInstanceKey() string
	GetCreateTime() time.Time
	GetState() sepc_element_types.LifecycleState
}
