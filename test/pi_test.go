package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/averyyan/bpmn-engine/bpmn/engine"
	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
	memory_engine "github.com/averyyan/bpmn-engine/memory"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestRun(t *testing.T) {
	state := memory_engine.New()
	state.TaskHandlerManager().RegisterServiceTaskHandler("test", func(activ engine_types.ActivatedActivity) error {
		fmt.Println("绑定程序运行")
		return activ.Complete()
	})
	pi, err := engine.CreateInstanceByFileAndRun(context.Background(), state, "cases/pi.bpmn", nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}
