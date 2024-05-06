package checkpoints

import (
	"github.com/TeaOSLab/EdgeNode/internal/waf/requests"
	"github.com/TeaOSLab/EdgeNode/internal/waf/utils"
	"github.com/iwind/TeaGo/maps"
)

type RequestHeaderMaxLengthCheckpoint struct {
	Checkpoint
}

func (this *RequestHeaderMaxLengthCheckpoint) RequestValue(req requests.Request, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	var maxLen int
	for _, v := range req.WAFRaw().Header {
		for _, subV := range v {
			var l = len(subV)
			if l > maxLen {
				maxLen = l
			}
		}
	}
	value = maxLen

	return
}

func (this *RequestHeaderMaxLengthCheckpoint) ResponseValue(req requests.Request, resp *requests.Response, param string, options maps.Map, ruleId int64) (value any, hasRequestBody bool, sysErr error, userErr error) {
	if this.IsRequest() {
		return this.RequestValue(req, param, options, ruleId)
	}
	return
}

func (this *RequestHeaderMaxLengthCheckpoint) CacheLife() utils.CacheLife {
	return utils.CacheShortLife
}
