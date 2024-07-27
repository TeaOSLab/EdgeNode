// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package readers

import (
	"io"
	"log"
)

type PrintReader struct {
	rawReader io.Reader
	tag       string
}

func NewPrintReader(rawReader io.Reader, tag string) io.Reader {
	return &PrintReader{
		rawReader: rawReader,
		tag:       tag,
	}
}

func (this *PrintReader) Read(p []byte) (n int, err error) {
	n, err = this.rawReader.Read(p)
	if n > 0 {
		log.Println("[" + this.tag + "]" + string(p[:n]))
	}
	return
}
