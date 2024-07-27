// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package nodes_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeCommon/pkg/nodeconfigs"
	"github.com/TeaOSLab/EdgeNode/internal/caches"
	"github.com/TeaOSLab/EdgeNode/internal/nodes"
)

func TestHTTPCacheTaskManager_Loop(t *testing.T) {
	// initialize cache policies
	config, err := nodeconfigs.SharedNodeConfig()
	if err != nil {
		t.Fatal(err)
	}
	caches.SharedManager.UpdatePolicies(config.HTTPCachePolicies)

	var manager = nodes.NewHTTPCacheTaskManager()
	err = manager.Loop()
	if err != nil {
		t.Fatal(err)
	}
}
