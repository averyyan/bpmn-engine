package flow

import (
	"html"
	"strings"

	"github.com/averyyan/bpmn-engine/bpmn/sepc/expression"
)

type TSequenceFlow struct {
	ID                  string                   `xml:"id,attr"`             // ID
	Name                string                   `xml:"name,attr"`           // 序列流名称
	SourceRef           string                   `xml:"sourceRef,attr"`      // 出发元素
	TargetRef           string                   `xml:"targetRef,attr"`      // 目标元素
	ConditionExpression []expression.TExpression `xml:"conditionExpression"` // 布尔表达式
}

func (flow *TSequenceFlow) HasConditionExpression() bool {
	return len(flow.ConditionExpression) == 1 && len(strings.TrimSpace(flow.GetConditionExpression())) > 0
}

func (flow *TSequenceFlow) GetConditionExpression() string {
	return html.UnescapeString(flow.ConditionExpression[0].Text)
}
