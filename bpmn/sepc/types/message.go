package sepc_types

type Message interface {
	GetID() string
	GetName() string
}
