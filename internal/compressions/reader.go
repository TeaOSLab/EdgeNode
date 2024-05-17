// Copyright 2021 GoEdge goedge.cdn@gmail.com. All rights reserved.

package compressions

import "io"

type Reader interface {
	Read(p []byte) (n int, err error)
	Reset(reader io.Reader) error
	RawClose() error
	Close() error
	IncreaseHit() uint32

	SetPool(pool *ReaderPool)
	ResetFinish()
}
