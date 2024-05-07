// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package caches

import "io"

type BaseReader struct {
	nextReader io.ReadCloser
}

// SetNextReader 设置下一个内容Reader
func (this *BaseReader) SetNextReader(nextReader io.ReadCloser) {
	this.nextReader = nextReader
}
