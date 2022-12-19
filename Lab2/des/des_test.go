package des

import "testing"

func TestDesEncryptDecrypt(t *testing.T) {
	key := []byte("12345678")
	text := []byte("learning")

	c, err := NewCipher(key)
	if err != nil {
		t.Fatalf("Expected no error but got %s", err)
	}

	encrypted := make([]byte, len(text))
	c.Encrypt(encrypted, text)

	if string(text) == string(encrypted) {
		t.Fatalf("Expected encrypted text to be different from %s", text)
	}

	decrypted := make([]byte, len(encrypted))
	c.Decrypt(decrypted, encrypted)

	if string(text) != string(decrypted) {
		t.Fatalf("Expected decrypted text to be %s text but got %s", text, decrypted)
	}
}
