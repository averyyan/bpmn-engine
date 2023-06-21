package sepc_types

// 任务元素接口
type ServiceTaskElement interface {
	// 基础元素接口
	BaseElement
	// 获取任务类型
	GetTaskDefinitionType() string
	// 获取外部数据
	GetProperties() []Propertie
}

// 任务元素接口
type UserTaskElement interface {
	// 基础元素接口
	BaseElement
	// 获取用户任务FromKey
	GetFormKey() string
	// 指定任务分配给特定用户 (用户任务类型)
	GetAssignmentAssignee() string
	// 指定任务分配给用户组 (用户任务类型)
	GetAssignmentCandidateUsers() []string
	// 指定任务分配给角色 (用户任务类型)
	GetAssignmentCandidateGroups() []string
}
