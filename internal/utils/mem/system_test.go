// Copyright 2021 GoEdge goedge.cdn@gmail.com. All rights reserved.

package memutils_test

import (
	"github.com/TeaOSLab/EdgeNode/internal/utils/mem"
	"testing"
)

func TestSystemMemoryGB(t *testing.T) {
	t.Log(memutils.SystemMemoryGB())
	t.Log(memutils.SystemMemoryGB())
	t.Log(memutils.SystemMemoryGB())
	t.Log(memutils.SystemMemoryBytes())
	t.Log(memutils.SystemMemoryBytes())
	t.Log(memutils.SystemMemoryBytes()>>30, "GB")
	t.Log("available:", memutils.AvailableMemoryGB())
}
