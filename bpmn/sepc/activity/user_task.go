package activity

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

// 表单任务
// task *TUserTask github.com/averyyan/bpmn-engine/bpmn/sepc/types.BaseElement
// task *TUserTask github.com/averyyan/bpmn-engine/bpmn/sepc/types.UserTask
type TUserTask struct {
	ID                   string                            `xml:"id,attr" json:"id"`                                                   // 元素ID
	Name                 string                            `xml:"name,attr" json:"name"`                                               // 元素名称
	Documentation        string                            `xml:"documentation,attr" json:"documentation"`                             // 元素说明
	IncomingAssociation  []string                          `xml:"incoming" json:"incoming"`                                            // 元素入Flow元素IDs
	OutgoingAssociation  []string                          `xml:"outgoing" json:"outgoing"`                                            // 元素出Flow元素IDs
	Properties           []*extensions.TPropertie          `xml:"extensionElements>properties>property" json:"properties"`             // 外部数据
	FormDefinition       *extensions.TFormDefinition       `xml:"extensionElements>formDefinition" json:"form_definition"`             // 表单信息
	AssignmentDefinition *extensions.TAssignmentDefinition `xml:"extensionElements>assignmentDefinition" json:"assignment_definition"` // 任务分配信息
}

func (task *TUserTask) GetID() string {
	return task.ID
}

func (task *TUserTask) GetName() string {
	return task.Name
}

func (task *TUserTask) GetType() sepc_element_types.ElementType {
	return sepc_element_types.UserTask
}

func (task *TUserTask) GetDocumentation() string {
	return task.Documentation
}

func (task *TUserTask) GetIncomingAssociation() []string {
	return task.IncomingAssociation
}

func (task *TUserTask) GetOutgoingAssociation() []string {
	return task.OutgoingAssociation
}

func (task *TUserTask) GetProperties() []sepc_types.Propertie {
	return extensions.Properties2Interface(task.Properties)
}

// 是否存在表单信息
func (task *TUserTask) HasFormDefinition() bool {
	return task.FormDefinition != nil
}

// 获取表单信息
func (task *TUserTask) GetFormDefinition() sepc_types.FormDefinition {
	return task.FormDefinition
}

// 是否存在分配信息
func (task *TUserTask) HasAssignmentDefinition() bool {
	return task.AssignmentDefinition != nil
}

// 获取分配信息
func (task *TUserTask) GetAssignmentDefinition() sepc_types.AssignmentDefinition {
	return task.AssignmentDefinition
}
