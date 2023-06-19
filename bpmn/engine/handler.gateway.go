package engine

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/flow"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

// 独占网关处理
func handleExclusiveGateway(
	flows []*flow.TSequenceFlow,
	variableContext map[string]interface{},
	element sepc_types.ExclusiveGateway,
) ([]*flow.TSequenceFlow, error) {
	return exclusivelyFilterByConditionExpression(flows, variableContext, element)
}

// 通过序列流上的表达式过滤
func exclusivelyFilterByConditionExpression(
	flows []*flow.TSequenceFlow,
	variableContext map[string]interface{},
	element sepc_types.ExclusiveGateway,
) ([]*flow.TSequenceFlow, error) {
	var ret []*flow.TSequenceFlow
	for _, flow := range flows {
		if flow.HasConditionExpression() {
			expression := flow.GetConditionExpression()
			out, err := evaluateExpression(expression, variableContext)
			if err != nil {
				return nil, err
			}
			if out == true {
				ret = append(ret, flow)
			}
		}
	}
	if len(ret) == 0 {
		ret = append(ret, findDefaultFlow(flows, element)...)
	}
	return ret, nil
}

// 找到默认序列流
func findDefaultFlow(flows []*flow.TSequenceFlow, element sepc_types.ExclusiveGateway) (ret []*flow.TSequenceFlow) {
	for _, flow := range flows {
		if flow.ID == element.GetDefault() {
			ret = append(ret, flow)
			break
		}
	}
	return ret
}
