package encrypt_test

import (
	"github.com/TeaOSLab/EdgeNode/internal/utils/encrypt"
	"testing"
)

func TestFindMethodInstance(t *testing.T) {
	t.Log(encrypt.NewMethodInstance("a", "b", ""))
	t.Log(encrypt.NewMethodInstance("aes-256-cfb", "123456", ""))
}
