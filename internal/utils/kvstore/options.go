// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package kvstore

import "github.com/cockroachdb/pebble"

var DefaultWriteOptions = &pebble.WriteOptions{
	Sync: false,
}

var DefaultWriteSyncOptions = &pebble.WriteOptions{
	Sync: true,
}
