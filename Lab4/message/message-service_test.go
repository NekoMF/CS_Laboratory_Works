package message

import (
	"testing"

	"github.com/Marcel-MD/cs-labs/hash/user"
)

func TestSignMessage(t *testing.T) {
	db := user.NewInMemoryDatabase()
	us := user.NewUserService(db)
	ms := NewMessageService()

	const (
		email    = "user@mail.com"
		password = "123456"
	)

	err := us.Register(email, password)
	if err != nil {
		t.Fatalf("Expected Register to return nil but got %v", err)
	}

	user, err := db.Get(email)
	if err != nil {
		t.Fatalf("Expected user to be registered but got error %v", err)
	}

	message := []byte("Hello World")

	signature, msgHashSum, err := ms.SignMessage(user.PrivateKey, message)
	if err != nil {
		t.Fatalf("Expected SignMessage to return nil but got %v", err)
	}

	err = ms.VerifyMessage(&user.PrivateKey.PublicKey, signature, msgHashSum)
	if err != nil {
		t.Fatalf("Expected VerifyMessage to return nil but got %v", err)
	}
}
