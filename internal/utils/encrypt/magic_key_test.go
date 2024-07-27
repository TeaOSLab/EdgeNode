package encrypt_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/utils/encrypt"
)

func TestMagicKeyEncode(t *testing.T) {
	var dst = encrypt.MagicKeyEncode([]byte("Hello,World"))
	t.Log("dst:", string(dst))

	var src = encrypt.MagicKeyDecode(dst)
	t.Log("src:", string(src))
}
