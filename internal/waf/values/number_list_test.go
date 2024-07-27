// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package values_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/waf/values"
	"github.com/iwind/TeaGo/assert"
)

func TestParseNumberList(t *testing.T) {
	var a = assert.NewAssertion(t)

	{
		var list = values.ParseNumberList("")
		a.IsFalse(list.Contains(123))
	}

	{
		var list = values.ParseNumberList(`123
456

789.1234`)
		a.IsTrue(list.Contains(123))
		a.IsFalse(list.Contains(0))
		a.IsFalse(list.Contains(789.123))
		a.IsTrue(list.Contains(789.1234))
	}
}
