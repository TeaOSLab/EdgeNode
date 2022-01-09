package nodes

import (
	"github.com/TeaOSLab/EdgeCommon/pkg/serverconfigs/firewallconfigs"
	"github.com/TeaOSLab/EdgeNode/internal/errors"
	"github.com/TeaOSLab/EdgeNode/internal/remotelogs"
	"github.com/TeaOSLab/EdgeNode/internal/waf"
	"strconv"
	"sync"
)

var sharedWAFManager = NewWAFManager()

// WAFManager WAF管理器
type WAFManager struct {
	mapping map[int64]*waf.WAF // policyId => WAF
	locker  sync.RWMutex
}

// NewWAFManager 获取新对象
func NewWAFManager() *WAFManager {
	return &WAFManager{
		mapping: map[int64]*waf.WAF{},
	}
}

// UpdatePolicies 更新策略
func (this *WAFManager) UpdatePolicies(policies []*firewallconfigs.HTTPFirewallPolicy) {
	this.locker.Lock()
	defer this.locker.Unlock()

	m := map[int64]*waf.WAF{}
	for _, p := range policies {
		w, err := this.convertWAF(p)
		if w != nil {
			m[p.Id] = w
		}
		if err != nil {
			remotelogs.Error("WAF", "initialize policy '"+strconv.FormatInt(p.Id, 10)+"' failed: "+err.Error())
			continue
		}
	}
	this.mapping = m
}

// FindWAF 查找WAF
func (this *WAFManager) FindWAF(policyId int64) *waf.WAF {
	this.locker.RLock()
	w, _ := this.mapping[policyId]
	this.locker.RUnlock()
	return w
}

// 将Policy转换为WAF
func (this *WAFManager) convertWAF(policy *firewallconfigs.HTTPFirewallPolicy) (*waf.WAF, error) {
	if policy == nil {
		return nil, errors.New("policy should not be nil")
	}
	if len(policy.Mode) == 0 {
		policy.Mode = firewallconfigs.FirewallModeDefend
	}
	w := &waf.WAF{
		Id:               policy.Id,
		IsOn:             policy.IsOn,
		Name:             policy.Name,
		Mode:             policy.Mode,
		UseLocalFirewall: policy.UseLocalFirewall,
	}

	// inbound
	if policy.Inbound != nil && policy.Inbound.IsOn {
		for _, group := range policy.Inbound.Groups {
			g := &waf.RuleGroup{
				Id:          group.Id,
				IsOn:        group.IsOn,
				Name:        group.Name,
				Description: group.Description,
				Code:        group.Code,
				IsInbound:   true,
			}

			// rule sets
			for _, set := range group.Sets {
				s := &waf.RuleSet{
					Id:          set.Id,
					Code:        set.Code,
					IsOn:        set.IsOn,
					Name:        set.Name,
					Description: set.Description,
					Connector:   set.Connector,
					IgnoreLocal: set.IgnoreLocal,
				}
				for _, a := range set.Actions {
					s.AddAction(a.Code, a.Options)
				}

				// rules
				for _, rule := range set.Rules {
					r := &waf.Rule{
						Description:       rule.Description,
						Param:             rule.Param,
						ParamFilters:      []*waf.ParamFilter{},
						Operator:          rule.Operator,
						Value:             rule.Value,
						IsCaseInsensitive: rule.IsCaseInsensitive,
						CheckpointOptions: rule.CheckpointOptions,
					}

					for _, paramFilter := range rule.ParamFilters {
						r.ParamFilters = append(r.ParamFilters, &waf.ParamFilter{
							Code:    paramFilter.Code,
							Options: paramFilter.Options,
						})
					}

					s.Rules = append(s.Rules, r)
				}

				g.RuleSets = append(g.RuleSets, s)
			}

			w.Inbound = append(w.Inbound, g)
		}
	}

	// outbound
	if policy.Outbound != nil && policy.Outbound.IsOn {
		for _, group := range policy.Outbound.Groups {
			g := &waf.RuleGroup{
				Id:          group.Id,
				IsOn:        group.IsOn,
				Name:        group.Name,
				Description: group.Description,
				Code:        group.Code,
				IsInbound:   true,
			}

			// rule sets
			for _, set := range group.Sets {
				s := &waf.RuleSet{
					Id:          set.Id,
					Code:        set.Code,
					IsOn:        set.IsOn,
					Name:        set.Name,
					Description: set.Description,
					Connector:   set.Connector,
					IgnoreLocal: set.IgnoreLocal,
				}

				for _, a := range set.Actions {
					s.AddAction(a.Code, a.Options)
				}

				// rules
				for _, rule := range set.Rules {
					r := &waf.Rule{
						Description:       rule.Description,
						Param:             rule.Param,
						Operator:          rule.Operator,
						Value:             rule.Value,
						IsCaseInsensitive: rule.IsCaseInsensitive,
						CheckpointOptions: rule.CheckpointOptions,
					}
					s.Rules = append(s.Rules, r)
				}

				g.RuleSets = append(g.RuleSets, s)
			}

			w.Outbound = append(w.Outbound, g)
		}
	}

	// action
	if policy.BlockOptions != nil {
		w.DefaultBlockAction = &waf.BlockAction{
			StatusCode: policy.BlockOptions.StatusCode,
			Body:       policy.BlockOptions.Body,
			URL:        policy.BlockOptions.URL,
			Timeout:    policy.BlockOptions.Timeout,
		}
	}

	errorList := w.Init()
	if len(errorList) > 0 {
		return w, errorList[0]
	}

	return w, nil
}
