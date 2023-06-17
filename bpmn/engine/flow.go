package engine

import "github.com/averyyan/bpmn-engine/bpmn/sepc/flow"

// 根据序列ID组找到对应的序列流
func findSequenceFlows(sequenceFlows []*flow.TSequenceFlow, ids []string) []*flow.TSequenceFlow {
	var ret []*flow.TSequenceFlow
	for _, flow := range sequenceFlows {
		for _, id := range ids {
			if id == flow.ID {
				ret = append(ret, flow)
			}
		}
	}
	return ret
}
