package process

import (
	"github.com/averyyan/bpmn-engine/bpmn/sepc/event"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/flow"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/gateway"
	"github.com/averyyan/bpmn-engine/bpmn/sepc/task"
)

// 流程内容
type TProcess struct {
	ID           string `xml:"id,attr"`           // ID
	Name         string `xml:"name,attr"`         // 名称
	IsExecutable bool   `xml:"isExecutable,attr"` // 是否启用
	// 随着流程元素增加添加
	StartEvents             []*event.TStartEvent             `xml:"startEvent"`             // 开始事件
	EndEvents               []*event.TEndEvent               `xml:"endEvent"`               // 结束事件
	IntermediateCatchEvents []*event.TIntermediateCatchEvent `xml:"intermediateCatchEvent"` // 中间事件
	SequenceFlows           []*flow.TSequenceFlow            `xml:"sequenceFlow"`           // 序列流
	ServiceTasks            []*task.TServiceTask             `xml:"serviceTask"`            // 服务任务
	ExclusiveGateways       []*gateway.TExclusiveGateway     `xml:"exclusiveGateway"`       // 独占网关
	ParallelGateways        []*gateway.TParallelGateway      `xml:"parallelGateway"`        // 并行网关
}
