// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package goman

import (
	"runtime"
	"sync"

	"github.com/TeaOSLab/EdgeNode/internal/utils/zero"
)

type TaskGroup struct {
	semi   chan zero.Zero
	wg     *sync.WaitGroup
	locker *sync.RWMutex
}

func NewTaskGroup() *TaskGroup {
	var concurrent = runtime.NumCPU()
	if concurrent <= 1 {
		concurrent = 2
	}
	return &TaskGroup{
		semi:   make(chan zero.Zero, concurrent),
		wg:     &sync.WaitGroup{},
		locker: &sync.RWMutex{},
	}
}

func (this *TaskGroup) Run(f func()) {
	this.wg.Add(1)
	go func() {
		defer this.wg.Done()

		this.semi <- zero.Zero{}

		f()

		<-this.semi
	}()
}

func (this *TaskGroup) Wait() {
	this.wg.Wait()
}

func (this *TaskGroup) Lock() {
	this.locker.Lock()
}

func (this *TaskGroup) Unlock() {
	this.locker.Unlock()
}
