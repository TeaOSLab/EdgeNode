// Copyright 2021 GoEdge goedge.cdn@gmail.com. All rights reserved.

package waf

import "github.com/iwind/TeaGo/maps"

type ActionConfig struct {
	Code    string   `yaml:"code" json:"code"`
	Options maps.Map `yaml:"options" json:"options"`
}
