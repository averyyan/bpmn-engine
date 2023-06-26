package sepc_types

// 重复流程元素
type CallActivity interface {
	HasCalledElement() bool
	GetCalledElement() CalledElement
}

// 执行流程信息
type CalledElement interface {
	GetProcessID() string
	GetPropagateAllChildVariables() bool
}

// 自动任务元素接口
type ServiceTask interface {
	// 获取任务定义
	HasTaskDefinition() bool
	// 获取任务类型
	GetTaskDefinition() TaskDefinition
}

// 任务定义
type TaskDefinition interface {
	// 任务类型
	GetTypeName() string
	// 任务重复次数
	GetRetries() string
}

// 任务元素接口
type UserTask interface {
	// 是否存在表单信息
	HasFormDefinition() bool
	// 获取表单信息
	GetFormDefinition() FormDefinition
	// 是否存在分配信息
	HasAssignmentDefinition() bool
	// 获取分配信息
	GetAssignmentDefinition() AssignmentDefinition
}

// 表单信息
type FormDefinition interface {
	GetFormKey() string
}

// 分配信息
type AssignmentDefinition interface {
}
