package expression

type TExpression struct {
	Text string `xml:",innerxml" bson:"text"`
}
