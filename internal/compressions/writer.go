// Copyright 2021 GoEdge goedge.cdn@gmail.com. All rights reserved.

package compressions

import "io"

type Writer interface {
	Write(p []byte) (int, error)
	Flush() error
	Reset(writer io.Writer)
	RawClose() error
	Close() error
	Level() int
	IncreaseHit() uint32

	SetPool(pool *WriterPool)
	ResetFinish()
}
