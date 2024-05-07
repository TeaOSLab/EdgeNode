// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cn .

package nodes

import (
	"github.com/TeaOSLab/EdgeNode/internal/caches"
	"github.com/iwind/TeaGo/types"
	"io"
	"net/http"
)

// HTTPRequestPartialReader 分区文件读取器
type HTTPRequestPartialReader struct {
	req    *HTTPRequest
	offset int64
	resp   *http.Response

	cacheReader caches.Reader
	cacheWriter caches.Writer
}

// NewHTTPRequestPartialReader 构建新的分区文件读取器
// req 当前请求
// offset 读取位置
// reader 当前缓存读取器
func NewHTTPRequestPartialReader(req *HTTPRequest, offset int64, reader caches.Reader) *HTTPRequestPartialReader {
	return &HTTPRequestPartialReader{
		req:         req,
		offset:      offset,
		cacheReader: reader,
	}
}

// 读取内容
func (this *HTTPRequestPartialReader) Read(p []byte) (n int, err error) {
	if this.resp == nil {
		_ = this.cacheReader.Close()

		this.req.RawReq.Header.Set("Range", "bytes="+types.String(this.offset)+"-")
		var resp = this.req.doReverseProxy(false)
		if resp == nil {
			err = io.ErrUnexpectedEOF
			return
		}

		this.resp = resp

		// 对比Content-MD5
		partialReader, ok := this.cacheReader.(*caches.PartialFileReader)
		if ok {
			if partialReader.Ranges().Version >= 2 && resp.Header.Get("Content-MD5") != partialReader.Ranges().ContentMD5 {
				err = io.ErrUnexpectedEOF

				var storage = this.req.writer.cacheStorage
				if storage != nil {
					_ = storage.Delete(this.req.cacheKey + caches.SuffixPartial)
				}

				return
			}
		}

		// 准备写入
		this.prepareCacheWriter()
	}

	n, err = this.resp.Body.Read(p)

	// 写入到缓存
	if n > 0 && this.cacheWriter != nil {
		_ = this.cacheWriter.WriteAt(this.offset, p[:n])
		this.offset += int64(n)
	}

	return
}

// Close 关闭读取器
func (this *HTTPRequestPartialReader) Close() error {
	if this.cacheWriter != nil {
		_ = this.cacheWriter.Close()
	}

	if this.resp != nil && this.resp.Body != nil {
		return this.resp.Body.Close()
	}

	return nil
}

// 准备缓存写入器
func (this *HTTPRequestPartialReader) prepareCacheWriter() {
	var storage = this.req.writer.cacheStorage
	if storage == nil {
		return
	}

	var cacheKey = this.req.cacheKey + caches.SuffixPartial
	writer, err := storage.OpenWriter(cacheKey, this.cacheReader.ExpiresAt(), this.cacheReader.Status(), int(this.cacheReader.HeaderSize()), this.cacheReader.BodySize(), -1, true)
	if err == nil {
		this.cacheWriter = writer
	}
}
