package engine

import (
	"context"

	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/flow"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

// 处理并行网关
func handleParallelGateway(
	ctx context.Context,
	state engine_types.Engine,
	pi engine_types.ProcessInstance,
	baseElement sepc_types.BaseElement,
) bool {
	allInboundsAreScheduled := true
	for _, inFlowId := range baseElement.GetIncomingAssociation() {
		allInboundsAreScheduled = state.ProcessInstanceManager().HasScheduledFlow(ctx, pi, inFlowId) && allInboundsAreScheduled
	}
	return allInboundsAreScheduled
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
