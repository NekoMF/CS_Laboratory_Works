package rc4

import "testing"

func TestXORKeyStream(t *testing.T) {
	key := []byte("secret")
	text := []byte("hello world")

	encrypted := make([]byte, len(text))
	XORKeyStream(encrypted, text, key)

	if string(text) == string(encrypted) {
		t.Fatalf("Expected encrypted text to be different from %s", text)
	}

	decrypted := make([]byte, len(encrypted))
	XORKeyStream(decrypted, encrypted, key)

	if string(text) != string(decrypted) {
		t.Fatalf("Expected decrypted text to be %s text but got %s", text, decrypted)
	}
}
