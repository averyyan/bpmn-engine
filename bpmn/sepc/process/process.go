package process

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/activity"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/event"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/extensions"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/flow"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/gateway"
)

// 流程内容
type TProcess struct {
	ID                      string                           `xml:"id,attr" json:"id"`                                       // ID
	Name                    string                           `xml:"name,attr" json:"name"`                                   // 名称
	Documentation           string                           `xml:"documentation,attr" json:"documentation"`                 // 流程说明
	IsExecutable            bool                             `xml:"isExecutable,attr" json:"is_executable"`                  // 是否启用
	StartEvents             []*event.TStartEvent             `xml:"startEvent" json:"start_events"`                          // 开始事件
	EndEvents               []*event.TEndEvent               `xml:"endEvent" json:"end_events"`                              // 结束事件
	IntermediateCatchEvents []*event.TIntermediateCatchEvent `xml:"intermediateCatchEvent" json:"intermediate_catch_events"` // 中间事件
	SequenceFlows           []*flow.TSequenceFlow            `xml:"sequenceFlow" json:"sequence_flows"`                      // 序列流
	ExclusiveGateways       []*gateway.TExclusiveGateway     `xml:"exclusiveGateway" json:"exclusive_gateways"`              // 独占网关
	ParallelGateways        []*gateway.TParallelGateway      `xml:"parallelGateway" json:"parallel_gateways"`                // 并行网关
	EventBasedGateways      []*gateway.TEventBasedGateway    `xml:"eventBasedGateway" json:"event_based_gateways"`           // 事件网关
	ServiceTasks            []*activity.TServiceTask         `xml:"serviceTask" json:"service_tasks"`                        // 服务任务
	CallActivitys           []*activity.TCallActivity        `xml:"callActivity" json:"call_activitys"`                      // 子流程
	Properties              []*extensions.TPropertie         `xml:"extensionElements>properties>property" json:"properties"` // 扩展数据
}
