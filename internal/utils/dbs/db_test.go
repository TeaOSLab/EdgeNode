// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package dbs_test

import (
	"net/url"
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/utils/dbs"
)

func TestParseDSN(t *testing.T) {
	var dsn = "file:/home/cache/p43/.indexes/db-3.db?cache=private&mode=ro&_journal_mode=WAL&_sync=" + dbs.SyncMode + "&_cache_size=88000"
	u, err := url.Parse(dsn)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(u.Path) // expect: :/home/cache/p43/.indexes/db-3.db
}
