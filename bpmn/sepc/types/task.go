package sepc_types

// 任务元素接口
type ServiceTask interface {
	// 获取任务定义
	HasTaskDefinition() bool
	// 获取任务类型
	GetTaskDefinition() TaskDefinition
	// 获取外部数据
	GetProperties() []Propertie
}

// 任务定义
type TaskDefinition interface {
	// 任务类型
	GetTypeName() string
	// 任务可重复次数
	GetRetries() string
}

// 任务元素接口
type UserTaskElement interface {
	// 获取用户任务FromKey
	GetFormKey() string
	// 指定任务分配给特定用户 (用户任务类型)
	GetAssignmentAssignee() string
	// 指定任务分配给用户组 (用户任务类型)
	GetAssignmentCandidateUsers() []string
	// 指定任务分配给角色 (用户任务类型)
	GetAssignmentCandidateGroups() []string
}
