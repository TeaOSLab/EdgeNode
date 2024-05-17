// Copyright 2022 GoEdge goedge.cdn@gmail.com. All rights reserved.

package stats

import (
	"github.com/mssola/useragent"
)

type UserAgentParserResult struct {
	OS             useragent.OSInfo
	BrowserName    string
	BrowserVersion string
	IsMobile       bool
}
