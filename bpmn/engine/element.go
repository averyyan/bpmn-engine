package engine

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/definitions"
	sepc_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types"
)

func makeElementMap(definitions *definitions.TDefinitions) map[string]sepc_types.BaseElement {
	elementMap := make(map[string]sepc_types.BaseElement)
	appender := func(element sepc_types.BaseElement) {
		elementMap[element.GetID()] = element
	}
	for _, startEvent := range definitions.Process.StartEvents {
		appender(startEvent)
	}
	for _, endEvent := range definitions.Process.EndEvents {
		appender(endEvent)
	}
	for _, task := range definitions.Process.ServiceTasks {
		appender(task)
	}
	for _, exclusiveGateway := range definitions.Process.ExclusiveGateways {
		appender(exclusiveGateway)
	}
	for _, parallelGateway := range definitions.Process.ParallelGateways {
		appender(parallelGateway)
	}
	for _, eventBasedGateway := range definitions.Process.EventBasedGateways {
		appender(eventBasedGateway)
	}
	for _, intermediateCatchEvent := range definitions.Process.IntermediateCatchEvents {
		appender(intermediateCatchEvent)
	}
	// for _, task := range definitions.Process.CallActivitys {
	// 	appender(task)
	// }
	// for _, task := range definitions.Process.UserTasks {
	// 	appender(task)
	// }
	return elementMap
}
