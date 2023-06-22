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

func TestExclusivelyGateway(t *testing.T) {
	v := 2
	state := memory_engine.New()
	fun := func(activ engine_types.ActivatedActivity) error {
		fmt.Println(activ.GetElement().GetName())
		return activ.Complete()
	}
	state.TaskHandlerManager().RegisterServiceTaskHandler("test0", fun)
	state.TaskHandlerManager().RegisterServiceTaskHandler("test1", fun)
	state.TaskHandlerManager().RegisterServiceTaskHandler("test2", fun)
	pi, err := engine.CreateInstanceByFileAndRun(context.Background(), state, "cases/exclusively-gateway.bpmn", map[string]any{"v": v})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}

func TestParallelGateway(t *testing.T) {
	v := 0
	state := memory_engine.New()
	fun := func(activ engine_types.ActivatedActivity) error {
		fmt.Println(activ.GetElement().GetName())
		return activ.Complete()
	}
	state.TaskHandlerManager().RegisterServiceTaskHandler("test0", fun)
	state.TaskHandlerManager().RegisterServiceTaskHandler("test1", fun)
	state.TaskHandlerManager().RegisterServiceTaskHandler("end1", func(activ engine_types.ActivatedActivity) error {
		return activ.Complete()
	})
	state.TaskHandlerManager().RegisterServiceTaskHandler("end2", func(activ engine_types.ActivatedActivity) error {
		return activ.Complete()
	})
	pi, err := engine.CreateInstanceByFileAndRun(context.Background(), state, "cases/parallel-gateway.bpmn", map[string]any{"v": v})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}

func TestMsgEvent(t *testing.T) {
	state := memory_engine.New()
	fun := func(activ engine_types.ActivatedActivity) error {
		fmt.Println(activ.GetElement().GetName())
		return activ.Complete()
	}
	state.TaskHandlerManager().RegisterServiceTaskHandler("test", fun)
	pi, err := engine.CreateInstanceByFileAndRun(context.Background(), state, "cases/msg-event.bpmn", nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Active))
	pi, err = engine.PublishEventForInstanceAndRun(context.Background(), state, pi.GetKey(), "test", nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}

func TestEventsBase(t *testing.T) {
	msg := "msg2"
	state := memory_engine.New()
	fun := func(activ engine_types.ActivatedActivity) error {
		fmt.Println(activ.GetElement().GetName())
		return activ.Complete()
	}
	state.TaskHandlerManager().RegisterServiceTaskHandler("test", fun)
	pi, err := engine.CreateInstanceByFileAndRun(context.Background(), state, "cases/events-base.bpmn", nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Active))
	pi, err = engine.PublishEventForInstanceAndRun(context.Background(), state, pi.GetKey(), msg, nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}
