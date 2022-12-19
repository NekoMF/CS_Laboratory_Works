package user

import (
	"testing"
)

func TestLoginRegister(t *testing.T) {

	db := NewInMemoryDatabase()

	us := NewUserService(db)

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

	if string(user.Password) == password {
		t.Fatalf("Expected password to be hashed but got %s", user.Password)
	}

	err = us.Login(email, password)
	if err != nil {
		t.Fatalf("Expected Login to return nil but got %v", err)
	}
}
