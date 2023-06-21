package extensions

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
