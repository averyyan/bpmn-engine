package message

type Message interface {
	GetID() string
	GetName() string
}

type TMessage struct {
	ID   string `xml:"id,attr"`   // 消息ID
	Name string `xml:"name,attr"` // 消息名称
}
