package extensions

type TPropertie struct {
	Name  string `xml:"name,attr" bson:"name"`
	Value string `xml:"value,attr" bson:"value"`
}
