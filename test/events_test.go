package test

import (
	"testing"

	bpmn_util "github.com/averyyan/bpmn-engine/bpmn/util"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestEvents(t *testing.T) {
	dfs, err := bpmn_util.LoadFromFile("cases/events.bpmn")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, len(dfs.Process.IntermediateCatchEvents), is.EqualTo(3))
	for _, ice := range dfs.Process.IntermediateCatchEvents {
		if ice.GetMessageEventDefinition() != nil {
			for _, msg := range dfs.Messages {
				if msg.ID == ice.GetMessageEventDefinition().GetMessageRef() {
					then.AssertThat(t, msg.Name, is.EqualTo("msg"))
				}
			}
		}
		if ice.GetTimerEventDefinition() != nil {
			if ice.GetTimerEventDefinition().GetTimeDate() != nil {
				t.Log(ice.GetTimerEventDefinition().GetTimeDate().GetText())
			}
			if ice.GetTimerEventDefinition().GetTimeDuration() != nil {
				t.Log(ice.GetTimerEventDefinition().GetTimeDuration().GetText())
			}
		}
	}
}
