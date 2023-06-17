package event

type TEndEvent struct {
	ID                  string   `xml:"id,attr"`   // 元素ID
	Name                string   `xml:"name,attr"` // 元素名称
	IncomingAssociation []string `xml:"incoming"`  // 元素入Flow元素IDs
	OutgoingAssociation []string `xml:"outgoing"`  // 元素出Flow元素IDs
}

func (endEvent *TEndEvent) GetID() string {
	return endEvent.ID
}

func (endEvent *TEndEvent) GetName() string {
	return endEvent.Name
}

func (endEvent *TEndEvent) GetIncomingAssociation() []string {
	return endEvent.IncomingAssociation
}

func (endEvent *TEndEvent) GetOutgoingAssociation() []string {
	return endEvent.OutgoingAssociation
}
