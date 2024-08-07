// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package fsutils_test

import (
	"testing"

	fsutils "github.com/TeaOSLab/EdgeNode/internal/utils/fs"
	"github.com/iwind/TeaGo/assert"
)

func TestFileFlags(t *testing.T) {
	var a = assert.NewAssertion(t)
	a.IsTrue(fsutils.FlagRead&fsutils.FlagRead == fsutils.FlagRead)
	a.IsTrue(fsutils.FlagWrite&fsutils.FlagWrite != fsutils.FlagRead)
	a.IsTrue((fsutils.FlagWrite|fsutils.FlagRead)&fsutils.FlagRead == fsutils.FlagRead)
}
