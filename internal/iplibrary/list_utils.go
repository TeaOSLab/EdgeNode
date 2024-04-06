// Copyright 2021 Liuxiangchao iwind.liu@gmail.com. All rights reserved.

package iplibrary

import (
	"bytes"
	"encoding/hex"
	"github.com/TeaOSLab/EdgeCommon/pkg/nodeconfigs"
	"github.com/iwind/TeaGo/Tea"
	"net"
)

// AllowIP 检查IP是否被允许访问
// 如果一个IP不在任何名单中，则允许访问
func AllowIP(ip string, serverId int64) (canGoNext bool, inAllowList bool, expiresAt int64) {
	if !Tea.IsTesting() { // 如果在测试环境，我们不加入一些白名单，以便于可以在本地和局域网正常测试
		// 放行lo
		if ip == "127.0.0.1" || ip == "::1" {
			return true, true, 0
		}

		// check node
		nodeConfig, err := nodeconfigs.SharedNodeConfig()
		if err == nil && nodeConfig.IPIsAutoAllowed(ip) {
			return true, true, 0
		}
	}

	var ipBytes = IPBytes(ip)
	if IsZero(ipBytes) {
		return false, false, 0
	}

	// check white lists
	if GlobalWhiteIPList.Contains(ipBytes) {
		return true, true, 0
	}

	if serverId > 0 {
		var list = SharedServerListManager.FindWhiteList(serverId, false)
		if list != nil && list.Contains(ipBytes) {
			return true, true, 0
		}
	}

	// check black lists
	expiresAt, ok := GlobalBlackIPList.ContainsExpires(ipBytes)
	if ok {
		return false, false, expiresAt
	}

	if serverId > 0 {
		var list = SharedServerListManager.FindBlackList(serverId, false)
		if list != nil {
			expiresAt, ok = list.ContainsExpires(ipBytes)
			if ok {
				return false, false, expiresAt
			}
		}
	}

	return true, false, 0
}

// IsInWhiteList 检查IP是否在白名单中
func IsInWhiteList(ip string) bool {
	var ipBytes = IPBytes(ip)
	if IsZero(ipBytes) {
		return false
	}

	// check white lists
	return GlobalWhiteIPList.Contains(ipBytes)
}

// AllowIPStrings 检查一组IP是否被允许访问
func AllowIPStrings(ipStrings []string, serverId int64) bool {
	if len(ipStrings) == 0 {
		return true
	}
	for _, ip := range ipStrings {
		isAllowed, _, _ := AllowIP(ip, serverId)
		if !isAllowed {
			return false
		}
	}
	return true
}

func IsZero(ipBytes []byte) bool {
	return len(ipBytes) == 0
}

func CompareBytes(b1 []byte, b2 []byte) int {
	var l1 = len(b1)
	var l2 = len(b2)
	if l1 < l2 {
		return -1
	}
	if l1 > l2 {
		return 1
	}
	return bytes.Compare(b1, b2)
}

func IPBytes(ip string) []byte {
	if len(ip) == 0 {
		return nil
	}

	var i = net.ParseIP(ip)
	if i == nil {
		return nil
	}

	var i4 = i.To4()
	if i4 != nil {
		return i4
	}

	return i.To16()
}

func ToHex(b []byte) string {
	if len(b) == 0 {
		return ""
	}

	return hex.EncodeToString(b)
}
