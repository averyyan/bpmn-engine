package sepc_types

type CallActivity interface {
	HasCalledElement() bool
	GetCalledElement() CalledElement
}

type CalledElement interface {
	GetProcessID() string
}
