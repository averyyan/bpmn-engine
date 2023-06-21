package task

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
	sepc_element_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/element"
)

type TServiceTask struct {
	ID                  string                     `xml:"id,attr"`                               // 元素ID
	Name                string                     `xml:"name,attr"`                             // 元素名称
	IncomingAssociation []string                   `xml:"incoming"`                              // 元素入Flow元素IDs
	OutgoingAssociation []string                   `xml:"outgoing"`                              // 元素出Flow元素IDs
	TaskDefinition      extensions.TTaskDefinition `xml:"extensionElements>taskDefinition"`      // Service Task定义
	Properties          []*extensions.TPropertie   `xml:"extensionElements>properties>property"` // 扩展数据
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

func (serviceTask *TServiceTask) GetTaskDefinitionType() string {
	return serviceTask.TaskDefinition.TypeName
}

func (serviceTask *TServiceTask) GetProperties() []sepc_types.Propertie {
	var properties []sepc_types.Propertie
	for _, v := range serviceTask.Properties {
		properties = append(properties, v)
	}
	return properties
}
