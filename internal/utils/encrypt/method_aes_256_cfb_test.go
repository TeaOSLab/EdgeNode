package encrypt_test

import (
	"testing"

	"github.com/TeaOSLab/EdgeNode/internal/utils/encrypt"
)

func TestAES256CFBMethod_Encrypt(t *testing.T) {
	method, err := encrypt.NewMethodInstance("aes-256-cfb", "abc", "123")
	if err != nil {
		t.Fatal(err)
	}
	var src = []byte("Hello, World")
	dst, err := method.Encrypt(src)
	if err != nil {
		t.Fatal(err)
	}
	dst = dst[:len(src)]
	t.Log("dst:", string(dst))

	src, err = method.Decrypt(dst)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("src:", string(src))
}

func TestAES256CFBMethod_Encrypt2(t *testing.T) {
	method, err := encrypt.NewMethodInstance("aes-256-cfb", "abc", "123")
	if err != nil {
		t.Fatal(err)
	}
	var src = []byte("Hello, World")
	dst, err := method.Encrypt(src)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("dst:", string(dst))

	src, err = method.Decrypt(dst)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("src:", string(src))
}
