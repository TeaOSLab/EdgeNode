// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package bfs_test

import (
	"math/rand"
	"strconv"
	"strings"
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/utils/bfs"
	"github.com/iwind/TeaGo/assert"
)

func TestCheckHash(t *testing.T) {
	var a = assert.NewAssertion(t)

	a.IsFalse(bfs.CheckHash("123456"))
	a.IsFalse(bfs.CheckHash(strings.Repeat("A", 32)))
	a.IsTrue(bfs.CheckHash(strings.Repeat("a", 32)))
	a.IsTrue(bfs.CheckHash(bfs.Hash("123456")))
}

func BenchmarkCheckHashErr(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = bfs.CheckHash(bfs.Hash(strconv.Itoa(rand.Int())))
	}
}
