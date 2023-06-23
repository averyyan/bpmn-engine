package task

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// serviceTask *TServiceTask github.com/averyyan/bpmn-engine/bpmn/sepc/types.BaseElement
// serviceTask *TServiceTask github.com/averyyan/bpmn-engine/bpmn/sepc/types.ServiceTask
type TServiceTask struct {
	ID                  string                      `xml:"id,attr"`                               // 元素ID
	Name                string                      `xml:"name,attr"`                             // 元素名称
	Documentation       string                      `xml:"documentation,attr"`                    // 元素说明
	IncomingAssociation []string                    `xml:"incoming"`                              // 元素入Flow元素IDs
	OutgoingAssociation []string                    `xml:"outgoing"`                              // 元素出Flow元素IDs
	Properties          []*extensions.TPropertie    `xml:"extensionElements>properties>property"` // 扩展数据
	TaskDefinition      *extensions.TTaskDefinition `xml:"extensionElements>taskDefinition"`      // Service Task定义
}

// 获取任务类型
func (serviceTask *TServiceTask) GetTaskDefinition() sepc_types.TaskDefinition {
	return serviceTask.TaskDefinition
}

func (serviceTask *TServiceTask) GetDocumentation() string {
	return serviceTask.Documentation
}

func (serviceTask *TServiceTask) GetID() string {
	return serviceTask.ID
}

func (serviceTask *TServiceTask) GetName() string {
	return serviceTask.Name
}

func (serviceTask *TServiceTask) GetIncomingAssociation() []string {
	return serviceTask.IncomingAssociation
}

func (serviceTask *TServiceTask) GetOutgoingAssociation() []string {
	return serviceTask.OutgoingAssociation
}

func (serviceTask *TServiceTask) GetType() sepc_element_types.ElementType {
	return sepc_element_types.ServiceTask
}

func (serviceTask *TServiceTask) HasTaskDefinition() bool {
	return serviceTask.TaskDefinition != nil
}

func (serviceTask *TServiceTask) GetProperties() []sepc_types.Propertie {
	return extensions.Properties2Interface(serviceTask.Properties)
}
