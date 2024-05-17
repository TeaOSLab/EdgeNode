// Copyright 2021 GoEdge goedge.cdn@gmail.com. All rights reserved.

package waf

import "github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"

type ActionCategory = string

const (
	ActionCategoryAllow  ActionCategory = firewallconfigs.HTTPFirewallActionCategoryAllow
	ActionCategoryBlock  ActionCategory = firewallconfigs.HTTPFirewallActionCategoryBlock
	ActionCategoryVerify ActionCategory = firewallconfigs.HTTPFirewallActionCategoryVerify
)
