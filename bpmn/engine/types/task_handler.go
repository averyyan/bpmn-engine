package engine_types

type TaskHandler func(activ ActivatedActivity) error

type TaskHandlerManager interface {
	// 注册 Servcie Task 元素处理函数
	RegisterServiceTaskHandler(taskType string, fun TaskHandler)
	// 获取 Servcie Task 元素处理函数
	GetServiceTaskHandler(taskType string) (TaskHandler, error)
}
