// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved.

package compressions

import (
	"io"

	teaconst "github.com/TeaOSLab/EdgeNode/internal/const"
)

var sharedZSTDReaderPool *ReaderPool

func init() {
	if !teaconst.IsMain {
		return
	}

	sharedZSTDReaderPool = NewReaderPool(CalculatePoolSize(), func(reader io.Reader) (Reader, error) {
		return newZSTDReader(reader)
	})
}
