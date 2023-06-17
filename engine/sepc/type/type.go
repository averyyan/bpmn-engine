package sepc_types

// 基础元素接口
type BaseElement interface {
	GetID() string                    // 获取元素ID
	GetName() string                  // 获取元素名称
	GetIncomingAssociation() []string // 获取元素入序列流
	GetOutgoingAssociation() []string // 获取元素出序列流
}

// 任务元素接口 待定
type TaskElement interface {
	BaseElement
	// GetInputMapping() []extensions.TIoMapping  // 获取进入规则
	// GetOutputMapping() []extensions.TIoMapping // 获取出规则
}

// 任务元素接口
type ServiceTaskElement interface {
	TaskElement
	GetTaskDefinitionType() string // 获取任务类型
}

// 任务元素接口
type UserTaskElement interface {
	TaskElement
	GetFormKey() string
	GetAssignmentAssignee() string          // 指定任务分配给特定用户 (用户任务类型)
	GetAssignmentCandidateUsers() []string  // 指定任务分配给用户组 (用户任务类型)
	GetAssignmentCandidateGroups() []string // 指定任务分配给角色 (用户任务类型)
}

// type ElementType string

// const (
// 	StartEvent             ElementType = "START_EVENT"              // 开始事件
// 	EndEvent               ElementType = "END_EVENT"                // 结束事件
// 	ServiceTask            ElementType = "SERVICE_TASK"             // 自动任务
// 	UserTask               ElementType = "USER_TASK"                // 人工任务
// 	ParallelGateway        ElementType = "PARALLEL_GATEWAY"         // 并行网关
// 	ExclusiveGateway       ElementType = "EXCLUSIVE_GATEWAY"        // 逻辑网关
// 	IntermediateCatchEvent ElementType = "INTERMEDIATE_CATCH_EVENT" // 消息捕获事件
// 	EventBasedGateway      ElementType = "EVENT_BASED_GATEWAY"      // 基于事件的网关
// 	SequenceFlow           ElementType = "SEQUENCE_FLOW"            // 箭头
// 	CallActivity           ElementType = "CALLACTIVITY"             // 重复子流程
// )
