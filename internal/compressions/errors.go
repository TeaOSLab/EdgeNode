// Copyright 2024 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .

package compressions

import "errors"

var ErrIsBusy = errors.New("the system is busy for compression")

func CanIgnore(err error) bool {
	if err == nil {
		return true
	}
	return errors.Is(err, ErrIsBusy)
}
