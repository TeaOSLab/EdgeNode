// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved.
//go:build !plus
// +build !plus

package nodes

import (
	"os"
)

func (this *HTTPWriter) canSendfile() (*os.File, bool) {
	return nil, false
}

func (this *HTTPWriter) checkPlanBandwidth(n int) {
	// stub
}
