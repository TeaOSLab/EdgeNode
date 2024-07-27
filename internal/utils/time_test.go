package utils_test

import (
	"testing"
	"time"

	"github.com/TeaOSLab/EdgeNode/internal/utils"
)

func TestGMTUnixTime(t *testing.T) {
	t.Log(utils.GMTUnixTime(time.Now().Unix()))
}

func TestGMTTime(t *testing.T) {
	t.Log(utils.GMTTime(time.Now()))
}
