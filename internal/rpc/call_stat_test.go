// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package rpc_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/rpc"
)

func TestNewCallStat(t *testing.T) {
	var stat = rpc.NewCallStat(10)
	stat.Add(true, 1)
	stat.Add(true, 2)
	stat.Add(true, 3)
	stat.Add(false, 4)
	stat.Add(true, 0)
	stat.Add(true, 1)
	t.Log(stat.Sum())
}
