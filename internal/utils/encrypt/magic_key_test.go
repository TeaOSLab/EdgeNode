package encrypt_test

import (
	"github.com/TeaOSLab/EdgeNode/internal/utils/encrypt"
	"testing"
)

func TestMagicKeyEncode(t *testing.T) {
	var dst = encrypt.MagicKeyEncode([]byte("Hello,World"))
	t.Log("dst:", string(dst))

	var src = encrypt.MagicKeyDecode(dst)
	t.Log("src:", string(src))
}
