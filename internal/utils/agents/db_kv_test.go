// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package agents_test

import (
	"strconv"
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/utils/agents"
)

func TestKVDB_InsertAgentIP(t *testing.T) {
	var db = agents.NewKVDB()
	err := db.Init()
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		_ = db.Flush()
	}()

	for i := 1; i <= 5; i++ {
		err = db.InsertAgentIP(int64(i), "192.168.2."+strconv.Itoa(i), "example")
		if err != nil {
			t.Fatal(err)
		}
	}
}

func TestKVDB_ListAgentIPs(t *testing.T) {
	var db = agents.NewKVDB()
	err := db.Init()
	if err != nil {
		t.Fatal(err)
	}

	const count = 10

	for {
		agentIPs, listErr := db.ListAgentIPs(0, count)
		if listErr != nil {
			t.Fatal(listErr)
		}
		t.Log("===")
		for _, agentIP := range agentIPs {
			t.Logf("%+v", agentIP)
		}

		if len(agentIPs) < count {
			break
		}
	}
}
