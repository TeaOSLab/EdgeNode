// Copyright 2023 GoEdge CDN goedge.cdn@gmail.com. All rights reserved. Official site: https://goedge.cloud .
//go:build !plus

package nodes

import (
	"errors"
	"net/http"

	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs"
)

func (this *HTTPRequest) doOSSOrigin(origin *serverconfigs.OriginConfig) (resp *http.Response, goNext bool, errorCode string, ossBucketName string, err error) {
	// stub
	return nil, false, "", "", errors.New("not implemented")
}
