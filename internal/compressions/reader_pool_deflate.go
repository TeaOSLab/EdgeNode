// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved.

package compressions

import (
	"io"

	teaconst "github.com/TeaOSLab/EdgeNode/internal/const"
)

var sharedDeflateReaderPool *ReaderPool

func init() {
	if !teaconst.IsMain {
		return
	}

	sharedDeflateReaderPool = NewReaderPool(CalculatePoolSize(), func(reader io.Reader) (Reader, error) {
		return newDeflateReader(reader)
	})
}
