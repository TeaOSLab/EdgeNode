// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package fsutils_test

import (
	"testing"

	fsutils "github.com/TeaOSLab/EdgeNode/internal/utils/fs"
)

func TestCheckDiskWritingSpeed(t *testing.T) {
	t.Log(fsutils.CheckDiskWritingSpeed())
}

func TestCheckDiskIsFast(t *testing.T) {
	t.Log(fsutils.CheckDiskIsFast())
}
