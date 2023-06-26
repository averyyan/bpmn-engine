package test

import (
	"context"
	"os"
	"testing"

	"github.com/averyyan/bpmn-engine/bpmn/engine"
	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	sepc_pi_types "github.com/averyyan/bpmn-engine/bpmn/sepc/types/process_instance"
	memory_process_instance "github.com/averyyan/bpmn-engine/memory/process_instance"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestServiceTask(t *testing.T) {
	processID := "Process_1vlirp4"
	ctx := context.Background()
	pi, err := engine.CreatePIByIDAndRun(ctx, processID, nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}

func TestExclusivelyGateway(t *testing.T) {
	v := 1
	processID := "Process_1ktaoh1"
	ctx := context.Background()
	pi, err := engine.CreatePIByIDAndRun(ctx, processID, map[string]any{"v": v})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}

func TestParallelGateway(t *testing.T) {
	v := 0
	processID := "Process_0ucn9gz"
	ctx := context.Background()
	pi, err := engine.CreatePIByIDAndRun(ctx, processID, map[string]any{"v": v})
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}

func TestMsgEvent(t *testing.T) {
	processID := "Process_1bzunxw"
	ctx := context.Background()
	pi, err := engine.CreatePIByIDAndRun(ctx, processID, nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Active))
	_, err = engine.PublishEventForInstanceContinue(ctx, pi.GetKey(), "test", nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}

func TestSavePI2JSON(t *testing.T) {
	processID := "Process_1bzunxw"
	ctx := context.Background()
	pi, err := engine.CreatePIByIDAndRun(ctx, processID, nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Active))
	raw, err := engine.SaveProcessInstance(pi)
	then.AssertThat(t, err, is.Nil())
	err = os.WriteFile("./json/pi.json", raw, 0777)
	then.AssertThat(t, err, is.Nil())
	pi2, err := engine.LoadProcessInstance[*memory_process_instance.ProcessInstance](raw)
	then.AssertThat(t, err, is.Nil())
	_, err = engine.PublishEventForInstanceContinue(ctx, pi2.GetKey(), "test", nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi2.GetState(), is.EqualTo(sepc_pi_types.Completed))
}

func TestEventsBase(t *testing.T) {
	msg := "msg1"
	processID := "Process_1b497xk"
	ctx := context.Background()
	pi, err := engine.CreatePIByIDAndRun(ctx, processID, nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Active))
	_, err = engine.PublishEventForInstanceContinue(ctx, pi.GetKey(), msg, nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}

func TestSavePI2JSON2(t *testing.T) {
	msg := "msg2"
	processID := "Process_1b497xk"
	ctx := context.Background()
	pi, err := engine.CreatePIByIDAndRun(ctx, processID, nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Active))
	raw, err := engine.SaveProcessInstance(pi)
	then.AssertThat(t, err, is.Nil())
	err = os.WriteFile("./json/pi.json", raw, 0777)
	then.AssertThat(t, err, is.Nil())
	pi2, err := engine.LoadProcessInstance[*memory_process_instance.ProcessInstance](raw)
	then.AssertThat(t, err, is.Nil())
	_, err = engine.PublishEventForInstanceContinue(ctx, pi2.GetKey(), msg, nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi2.GetState(), is.EqualTo(sepc_pi_types.Completed))
}

func TestCallActivity(t *testing.T) {
	processID := "Process_1g6e2gd"
	ctx := context.Background()
	pi, err := engine.CreatePIByIDAndRun(ctx, processID, nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Active))
	call, err := pi.GetElementManager().FindOneCallActivity(ctx, "Activity_1x8w1ty")
	then.AssertThat(t, err, is.Nil())
	childPI, err := engine.PublishEventForInstanceContinue(ctx, call.(engine_types.CallActivity).GetChildPIKey(), "msg1", nil)
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, childPI.GetState(), is.EqualTo(sepc_pi_types.Completed))
	then.AssertThat(t, pi.GetState(), is.EqualTo(sepc_pi_types.Completed))
}
