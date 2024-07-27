// Copyright 2023 GoEdge goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package maputils_test

import (
	"testing"

	maputils "github.com/TeaOSLab/EdgeNode/internal/utils/maps"
)

func TestNewFixedMap(t *testing.T) {
	var m = maputils.NewFixedMap[string, int](3)
	m.Put("a", 1)
	t.Log(m.RawMap())
	t.Log(m.Get("a"))
	t.Log(m.Get("b"))

	m.Put("b", 2)
	m.Put("c", 3)
	t.Log(m.RawMap(), m.Keys())

	m.Put("d", 4)
	t.Log(m.RawMap(), m.Keys())

	m.Put("b", 200)
	t.Log(m.RawMap(), m.Keys())
}
