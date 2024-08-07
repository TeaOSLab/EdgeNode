// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package utils_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/utils"
	"github.com/iwind/TeaGo/assert"
)

func TestIsCommonFileExtension(t *testing.T) {
	var a = assert.NewAssertion(t)

	a.IsTrue(utils.IsCommonFileExtension(".jpg"))
	a.IsTrue(utils.IsCommonFileExtension("png"))
	a.IsTrue(utils.IsCommonFileExtension("PNG"))
	a.IsTrue(utils.IsCommonFileExtension(".PNG"))
	a.IsTrue(utils.IsCommonFileExtension("Png"))
	a.IsFalse(utils.IsCommonFileExtension("zip"))
}
