// Copyright 2021 GoEdge goedge.cdn@gmail.com. All rights reserved.
//go:build !plus

package nodes

// 检查套餐
func (this *HTTPRequest) doPlanBefore() (blocked bool) {
	// stub
	return false
}
