package extensions

// td *TTaskDefinition github.com/averyyan/bpmn-engine/bpmn/sepc/types.TaskDefinition
type TTaskDefinition struct {
	TypeName string `xml:"type,attr"`
	Retries  string `xml:"retries,attr"`
}

func (td *TTaskDefinition) GetTypeName() string {
	return td.TypeName
}

func (td *TTaskDefinition) GetRetries() string {
	return td.Retries
}
