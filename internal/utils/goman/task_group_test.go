// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package goman_test

import (
	"runtime"
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/utils/goman"
)

func TestNewTaskGroup(t *testing.T) {
	var group = goman.NewTaskGroup()
	var m = map[int]bool{}

	for i := 0; i < runtime.NumCPU()*2; i++ {
		var index = i
		group.Run(func() {
			t.Log("task", index)

			group.Lock()
			_, ok := m[index]
			if ok {
				t.Error("duplicated:", index)
			}
			m[index] = true
			group.Unlock()
		})
	}
	group.Wait()
}
