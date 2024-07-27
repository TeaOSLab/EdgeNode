// Copyright 2021 GoEdge goedge.cdn@gmail.com. All rights reserved.

package nodes

import (
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/utils/testutils"
	_ "github.com/iwind/TeaGo/bootstrap"
)

func TestUpgradeManager_install(t *testing.T) {
	if !testutils.IsSingleTesting() {
		return
	}

	err := NewUpgradeManager().install()
	if err != nil {
		t.Fatal(err)
	}
	t.Log("ok")
}
