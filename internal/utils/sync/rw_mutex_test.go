// Copyright 2023 GoEdge goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package syncutils_test

import (
	"runtime"
	"sync"
	"testing"
	"time"

	syncutils "github.com/TeaOSLab/EdgeNode/internal/utils/sync"
)

func TestNewRWMutex(t *testing.T) {
	var locker = syncutils.NewRWMutex(runtime.NumCPU())
	locker.Lock(1)
	t.Log(locker.TryLock(1))
	locker.Unlock(1)
	t.Log(locker.TryLock(1))
}

func BenchmarkRWMutex_Lock(b *testing.B) {
	var locker = syncutils.NewRWMutex(runtime.NumCPU())

	b.RunParallel(func(pb *testing.PB) {
		var i = 0
		for pb.Next() {
			i++
			locker.Lock(i)
			time.Sleep(1 * time.Millisecond)
			locker.Unlock(i)
		}
	})
}

func BenchmarkRWMutex_Lock2(b *testing.B) {
	var locker = &sync.Mutex{}

	b.RunParallel(func(pb *testing.PB) {
		var i = 0
		for pb.Next() {
			i++
			locker.Lock()
			time.Sleep(1 * time.Millisecond)
			locker.Unlock()
		}
	})
}
