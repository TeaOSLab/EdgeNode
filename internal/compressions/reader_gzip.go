// Copyright 2021 GoEdge goedge.cdn@gmail.com. All rights reserved.

package compressions

import (
	"io"

	"github.com/klauspost/compress/gzip"
)

type GzipReader struct {
	BaseReader

	reader *gzip.Reader
}

func NewGzipReader(reader io.Reader) (Reader, error) {
	return sharedGzipReaderPool.Get(reader)
}

func newGzipReader(reader io.Reader) (Reader, error) {
	r, err := gzip.NewReader(reader)
	if err != nil {
		return nil, err
	}
	return &GzipReader{
		reader: r,
	}, nil
}

func (this *GzipReader) Read(p []byte) (n int, err error) {
	return this.reader.Read(p)
}

func (this *GzipReader) Reset(reader io.Reader) error {
	return this.reader.Reset(reader)
}

func (this *GzipReader) RawClose() error {
	return this.reader.Close()
}

func (this *GzipReader) Close() error {
	return this.Finish(this)
}
