package waf

import (
	"net/http"

	"github.com/TeaOSLab/EdgeNode/internal/waf/requests"
)

type TagAction struct {
	BaseAction

	Tags []string `yaml:"tags" json:"tags"`
}

func (this *TagAction) Init(waf *WAF) error {
	return nil
}

func (this *TagAction) Code() string {
	return ActionTag
}

func (this *TagAction) IsAttack() bool {
	return false
}

func (this *TagAction) WillChange() bool {
	return false
}

func (this *TagAction) Perform(waf *WAF, group *RuleGroup, set *RuleSet, request requests.Request, writer http.ResponseWriter) PerformResult {
	return PerformResult{
		ContinueRequest: true,
	}
}
