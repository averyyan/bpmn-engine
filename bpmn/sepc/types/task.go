package sepc_types

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
