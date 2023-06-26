package extensions

import (
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

// propertie *TPropertie github.com/averyyan/bpmn-engine/bpmn/sepc/types.Propertie
type TPropertie struct {
	Name  string `xml:"name,attr" json:"name"`
	Value string `xml:"value,attr" json:"value"`
}

// 获取扩展数据名称
func (propertie *TPropertie) GetName() string {
	return propertie.Name
}

// 获取扩展数据值
func (propertie *TPropertie) GetValue() string {
	return propertie.Value
}

// 将扩展数据值转换为接口
func Properties2Interface(properties []*TPropertie) []sepc_types.Propertie {
	var iproperties []sepc_types.Propertie
	for _, propertie := range properties {
		iproperties = append(iproperties, propertie)
	}
	return iproperties
}
