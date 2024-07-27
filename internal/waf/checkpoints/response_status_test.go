package checkpoints

import (
	"net/http"
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/waf/requests"
)

func TestResponseStatusCheckpoint_ResponseValue(t *testing.T) {
	resp := requests.NewResponse(new(http.Response))
	resp.StatusCode = 200

	checkpoint := new(ResponseStatusCheckpoint)
	t.Log(checkpoint.ResponseValue(nil, resp, "", nil, 1))
}
