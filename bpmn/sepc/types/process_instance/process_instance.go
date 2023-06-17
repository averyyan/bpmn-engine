package sepc_pi_types

type State string

const (
	Ready     State = "READY"     // 实例开始
	Active    State = "ACTIVE"    // 实例运行中
	Completed State = "COMPLETED" // 实例已完成
	Failed    State = "FAILED"    // 实例失败
)
