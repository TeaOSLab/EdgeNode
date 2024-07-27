package nodes

import (
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/utils/testutils"
)

func TestAPIStream_Start(t *testing.T) {
	if !testutils.IsSingleTesting() {
		return
	}

	apiStream := NewAPIStream()
	apiStream.Start()
}
