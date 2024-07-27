// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package nodes

type LingerConn interface {
	SetLinger(sec int) error
}
