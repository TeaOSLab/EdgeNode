// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package kvstore

type TableInterface interface {
	Name() string
	SetNamespace(namespace []byte)
	SetDB(db *DB)
	Close() error
}
