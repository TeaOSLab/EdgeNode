// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package fsutils_test

import (
	"testing"
	"time"

	fsutils "github.com/TeaOSLab/EdgeNode/internal/utils/fs"
)

func TestWaitLoad(t *testing.T) {
	fsutils.WaitLoad(100, 5, 1*time.Minute)
}
