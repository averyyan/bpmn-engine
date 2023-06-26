package event

// med *TMessageEventDefinition github.com/averyyan/bpmn-engine/bpmn/sepc/types.MessageEventDefinition
type TMessageEventDefinition struct {
	ID         string `xml:"id,attr" json:"id"`                  // 消息中间件ID
	MessageRef string `xml:"messageRef,attr" json:"message_ref"` // 消息中间件匹配消息订阅ID
}

func (med *TMessageEventDefinition) GetID() string {
	return med.ID
}

func (med *TMessageEventDefinition) GetMessageRef() string {
	return med.MessageRef
}
