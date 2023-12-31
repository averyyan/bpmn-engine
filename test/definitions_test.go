package test

import (
	"testing"

	bpmn_util "github.com/averyyan/bpmn-engine/bpmn/util"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
)

func TestDefinitions(t *testing.T) {
	dfs, err := bpmn_util.LoadFromFile("cases/definitions.bpmn")
	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, dfs.ID, is.EqualTo("Definitions_11zanqj"))
	then.AssertThat(t, dfs.Process.Name, is.EqualTo("全局解析测试"))
	then.AssertThat(t, len(dfs.Messages), is.EqualTo(1))
}
