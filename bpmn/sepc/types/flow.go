package sepc_types

type SequenceFlow interface {
	// 是否存在表达式
	HasConditionExpression() bool
	// 获取表达式
	GetConditionExpression() string
}
