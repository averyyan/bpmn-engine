package message

// msg *TMessage github.com/averyyan/bpmn-engine/bpmn/sepc/types.Message
type TMessage struct {
	ID   string `xml:"id,attr"`   // 消息ID
	Name string `xml:"name,attr"` // 消息名称
}

func (msg *TMessage) GetID() string {
	return msg.ID
}

func (msg *TMessage) GetName() string {
	return msg.Name
}
