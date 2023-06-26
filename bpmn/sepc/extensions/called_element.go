package extensions

// ce *TCalledElement github.com/averyyan/bpmn-engine/bpmn/sepc/types.CalledElement
type TCalledElement struct {
	ProcessID                  string `xml:"processId,attr" json:"process_id"`
	PropagateAllChildVariables bool   `xml:"propagateAllChildVariables,attr" json:"propagate_all_child_variables"`
}

func (ce *TCalledElement) GetProcessID() string {
	return ce.ProcessID
}

func (ce *TCalledElement) GetPropagateAllChildVariables() bool {
	return ce.PropagateAllChildVariables
}
