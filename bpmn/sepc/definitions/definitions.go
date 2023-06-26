package definitions

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/message"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/process"
)

// 流程信息文件
type TDefinitions struct {
	ID       string              `xml:"id,attr" json:"id"`       // ID
	Name     string              `xml:"name,attr" json:"name"`   // 流程文件名称
	Process  *process.TProcess   `xml:"process" json:"process"`  // 流程内容
	Messages []*message.TMessage `xml:"message" json:"messages"` // 消息定义
}
