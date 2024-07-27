package nodes

import (
	"testing"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

func TestNodeStatusExecutor_CPU(t *testing.T) {
	countLogicCPU, err := cpu.Counts(true)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("logic count:", countLogicCPU)

	countPhysicalCPU, err := cpu.Counts(false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("physical count:", countPhysicalCPU)

	percents, err := cpu.Percent(100*time.Millisecond, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(percents)
}
