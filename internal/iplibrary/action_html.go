package iplibrary

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/rpc/pb"
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"net/http"
)

// HTML动作
type HTMLAction struct {
	BaseAction

	config *firewallconfigs.FirewallActionHTMLConfig
}

// 获取新对象
func NewHTMLAction() *HTMLAction {
	return &HTMLAction{}
}

// 初始化
func (this *HTMLAction) Init(config *firewallconfigs.FirewallActionConfig) error {
	this.config = &firewallconfigs.FirewallActionHTMLConfig{}
	err := this.convertParams(config.Params, this.config)
	if err != nil {
		return err
	}
	return nil
}

// 添加
func (this *HTMLAction) AddItem(listType IPListType, item *pb.IPItem) error {
	return nil
}

// 删除
func (this *HTMLAction) DeleteItem(listType IPListType, item *pb.IPItem) error {
	return nil
}

// 关闭
func (this *HTMLAction) Close() error {
	return nil
}

// 处理HTTP请求
func (this *HTMLAction) DoHTTP(req *http.Request, resp http.ResponseWriter) (goNext bool, err error) {
	if this.config == nil {
		goNext = true
		return
	}
	resp.WriteHeader(http.StatusForbidden) // TODO改成可以配置
	_, _ = resp.Write([]byte(this.config.Content))
	return false, nil
}
