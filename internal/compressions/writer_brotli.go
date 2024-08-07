// Copyright 2021 GoEdge goedge.cdn@gmail.com. All rights reserved.
//go:build !plus || !linux

package compressions

import (
	"io"

	"github.com/andybalholm/brotli"
)

type BrotliWriter struct {
	BaseWriter

	writer *brotli.Writer
	level  int
}

func NewBrotliWriter(writer io.Writer, level int) (Writer, error) {
	return sharedBrotliWriterPool.Get(writer, level)
}

func newBrotliWriter(writer io.Writer) (*BrotliWriter, error) {
	var level = GenerateCompressLevel(brotli.BestSpeed, brotli.BestCompression)
	return &BrotliWriter{
		writer: brotli.NewWriterOptions(writer, brotli.WriterOptions{
			Quality: level,
			LGWin:   14, // TODO 在全局设置里可以设置此值
		}),
		level: level,
	}, nil
}

func (this *BrotliWriter) Write(p []byte) (int, error) {
	return this.writer.Write(p)
}

func (this *BrotliWriter) Flush() error {
	return this.writer.Flush()
}

func (this *BrotliWriter) Reset(newWriter io.Writer) {
	this.writer.Reset(newWriter)
}

func (this *BrotliWriter) RawClose() error {
	return this.writer.Close()
}

func (this *BrotliWriter) Close() error {
	return this.Finish(this)
}

func (this *BrotliWriter) Level() int {
	return this.level
}
