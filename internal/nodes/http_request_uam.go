// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved.
//go:build !plus
// +build !plus

package nodes

func (this *HTTPRequest) isUAMRequest() bool {
	// stub
	return false
}

// UAM
func (this *HTTPRequest) doUAM() (block bool) {
	// stub
	return false
}
