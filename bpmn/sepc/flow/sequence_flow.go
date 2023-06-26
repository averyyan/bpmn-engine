package flow

import (
	"html"
	"strings"

	"github.com/averyyan/bpmn-engine/bpmn/sepc/expression"
)

// flow *TSequenceFlow github.com/averyyan/bpmn-engine/bpmn/sepc/types.SequenceFlow
type TSequenceFlow struct {
	ID                  string                   `xml:"id,attr" json:"id"`                               // ID
	Name                string                   `xml:"name,attr" json:"name"`                           // 序列流名称
	SourceRef           string                   `xml:"sourceRef,attr" json:"source_ref"`                // 出发元素
	TargetRef           string                   `xml:"targetRef,attr" json:"target_ref"`                // 目标元素
	ConditionExpression []expression.TExpression `xml:"conditionExpression" json:"condition_expression"` // 布尔表达式
}

// 是否存在表达式
func (flow *TSequenceFlow) HasConditionExpression() bool {
	return len(flow.ConditionExpression) == 1 && len(strings.TrimSpace(flow.GetConditionExpression())) > 0
}

// 获取表达式
func (flow *TSequenceFlow) GetConditionExpression() string {
	return html.UnescapeString(flow.ConditionExpression[0].Text)
}
