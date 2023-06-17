package test

import (
	"context"
	"testing"

	"github.com/averyyan/bpmn-engine/bpmn/engine"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
	memory_engine "github.com/averyyan/bpmn-engine/memory"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestRun(t *testing.T) {
	state := memory_engine.New()
	pi, err := engine.CreateInstanceByFileAndRun(context.Background(), state, "cases/pi.bpmn", nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}
