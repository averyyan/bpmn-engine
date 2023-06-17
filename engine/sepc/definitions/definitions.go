package definitions

import (
	"github.com/averyyan/bpmn-engine/engine/sepc/message"
	"github.com/averyyan/bpmn-engine/engine/sepc/process"
)

type TDefinitions struct {
	ID       string              `xml:"id,attr"`   // ID
	Name     string              `xml:"name,attr"` // 流程文件名称
	Process  *process.TProcess   `xml:"process"`   // 流程内容
	Messages []*message.TMessage `xml:"message"`   // 消息定义详情
}
