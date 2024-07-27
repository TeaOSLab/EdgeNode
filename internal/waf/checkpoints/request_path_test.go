package checkpoints

import (
	"net/http"
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/waf/requests"
)

func TestRequestPathCheckpoint_RequestValue(t *testing.T) {
	rawReq, err := http.NewRequest(http.MethodGet, "http://teaos.cn/index?name=lu", nil)
	if err != nil {
		t.Fatal(err)
	}

	req := requests.NewTestRequest(rawReq)
	checkpoint := new(RequestPathCheckpoint)
	t.Log(checkpoint.RequestValue(req, "", nil, 1))
}
