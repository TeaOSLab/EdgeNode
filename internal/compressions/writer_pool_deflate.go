// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved.

package compressions

import (
	"io"

	teaconst "github.com/TeaOSLab/EdgeNode/internal/const"
)

var sharedDeflateWriterPool *WriterPool

func init() {
	if !teaconst.IsMain {
		return
	}

	sharedDeflateWriterPool = NewWriterPool(CalculatePoolSize(), func(writer io.Writer, level int) (Writer, error) {
		return newDeflateWriter(writer)
	})
}
