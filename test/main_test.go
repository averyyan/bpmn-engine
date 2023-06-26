package test

import (
	"embed"
	"fmt"
	"os"
	"testing"

	"github.com/averyyan/bpmn-engine/bpmn/engine"
	engine_types "github.com/averyyan/bpmn-engine/bpmn/engine/types"
	memory_process_instance "github.com/averyyan/bpmn-engine/memory/process_instance"
)

//go:embed cases
var processes embed.FS

func TestMain(m *testing.M) {
	engine.LoadFromEmbed(&processes)
	engine.RegisterPICreateHandler(memory_process_instance.New)
	fun := func(activ engine_types.ActivatedActivity) error {
		fmt.Println(activ.GetElement().GetName())
		return nil
	}
	engine.RegisterServiceTaskHandler("test", fun)
	engine.RegisterServiceTaskHandler("test0", fun)
	engine.RegisterServiceTaskHandler("test1", fun)
	engine.RegisterServiceTaskHandler("test2", fun)
	engine.RegisterServiceTaskHandler("end1", fun)
	engine.RegisterServiceTaskHandler("end2", fun)
	os.Exit(m.Run())
}
