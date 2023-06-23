package extensions

// calledElement *TCalledElement github.com/averyyan/bpmn-engine/bpmn/sepc/types.CalledElement
type TCalledElement struct {
	ProcessID string `xml:"processId,attr"`
}

func (calledElement *TCalledElement) GetProcessID() string {
	return calledElement.ProcessID
}
