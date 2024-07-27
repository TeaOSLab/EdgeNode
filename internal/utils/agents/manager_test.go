// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package agents_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/utils/agents"
	"github.com/TeaOSLab/EdgeNode/internal/utils/testutils"
	"github.com/iwind/TeaGo/Tea"
	_ "github.com/iwind/TeaGo/bootstrap"
)

func TestNewManager(t *testing.T) {
	if !testutils.IsSingleTesting() {
		return
	}

	var db = agents.NewSQLiteDB(Tea.Root + "/data/agents.db")
	err := db.Init()
	if err != nil {
		t.Fatal(err)
	}

	var manager = agents.NewManager()
	manager.SetDB(db)
	err = manager.Load()
	if err != nil {
		t.Fatal(err)
	}

	_, err = manager.Loop()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(manager.LookupIP("192.168.3.100"))   // not found
	t.Log(manager.LookupIP("66.249.79.25"))    // google
	t.Log(manager.ContainsIP("66.249.79.25"))  // true
	t.Log(manager.ContainsIP("66.249.79.255")) // not found
}
