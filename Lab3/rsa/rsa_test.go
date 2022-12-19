package rsa

import (
	"crypto/rand"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	priv, err := GenerateKey(rand.Reader, 1024)
	if err != nil {
		t.Fatal(err)
	}

	msg := []byte("hello world")
	c := Encrypt(&priv.PublicKey, msg)
	m := Decrypt(priv, c)

	if string(m) != string(msg) {
		t.Fatalf("expected %s, got %s", msg, m)
	}
}
