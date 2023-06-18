package extensions

type TTaskDefinition struct {
	TypeName string `xml:"type,attr" bson:"name"`
	Retries  string `xml:"retries,attr" bson:"retries"`
}
