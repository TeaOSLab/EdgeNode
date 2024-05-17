// Copyright 2021 GoEdge goedge.cdn@gmail.com. All rights reserved.

package goman_test

import (
	"github.com/TeaOSLab/EdgeNode/internal/utils/goman"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	goman.New(func() {
		t.Log("Hello")

		t.Log(goman.List())
	})

	time.Sleep(1 * time.Second)
	t.Log(goman.List())

	time.Sleep(1 * time.Second)
}

func TestNewWithArgs(t *testing.T) {
	goman.NewWithArgs(func(args ...interface{}) {
		t.Log(args[0], args[1])
	}, 1, 2)
	time.Sleep(1 * time.Second)
}
