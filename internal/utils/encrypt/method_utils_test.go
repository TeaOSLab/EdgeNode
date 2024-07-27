package encrypt_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/utils/encrypt"
)

func TestFindMethodInstance(t *testing.T) {
	t.Log(encrypt.NewMethodInstance("a", "b", ""))
	t.Log(encrypt.NewMethodInstance("aes-256-cfb", "123456", ""))
}
