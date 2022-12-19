package user

import (
	"crypto/rand"
	"crypto/rsa"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(email, password string) error
	Login(email, password string) error
}

type userService struct {
	db Database
}

func NewUserService(database Database) UserService {
	return &userService{db: database}
}

func (s *userService) Register(email, password string) error {

	_, err := s.db.Get(email)
	if err == nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	privKey, err := rsa.GenerateKey(rand.Reader, 1028)
	if err != nil {
		return err
	}

	user := User{
		Email:      email,
		Password:   hashedPassword,
		PrivateKey: privKey,
	}

	return s.db.Set(email, user)
}

func (s *userService) Login(email, password string) error {
	user, err := s.db.Get(email)
	if err != nil {
		return err
	}

	return bcrypt.CompareHashAndPassword(user.Password, []byte(password))
}
