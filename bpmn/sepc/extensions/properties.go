package extensions

import (
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

type TPropertie struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

func (propertie *TPropertie) GetName() string {
	return propertie.Name
}

func (propertie *TPropertie) GetValue() string {
	return propertie.Value
}

func Properties2Interface(properties []*TPropertie) []sepc_types.Propertie {
	var iproperties []sepc_types.Propertie
	for _, propertie := range properties {
		iproperties = append(iproperties, propertie)
	}
	return iproperties
}
