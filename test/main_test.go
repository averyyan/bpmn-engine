package test

import (
	"os"
	"testing"

	bpmn_util "github.com/averyyan/bpmn-engine/bpmn/util"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestXXX(t *testing.T) {
	dfs, _ := bpmn_util.LoadFromFile("cases/subprocess.bpmn")
	t.Log(dfs.Process.Properties)
}
