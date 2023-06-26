package extensions

// fd *TFormDefinition github.com/averyyan/bpmn-engine/bpmn/sepc/types.FormDefinition
type TFormDefinition struct {
	FormKey string `xml:"formKey,attr" json:"form_key"`
}

func (fd *TFormDefinition) GetFormKey() string {
	return fd.FormKey
}
