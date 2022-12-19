package api

import (
	"errors"

	"github.com/Marcel-MD/cs-labs/api/token"

	"github.com/Marcel-MD/cs-labs/api/dto"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Get(email string) (User, error)
	GetAll() ([]User, error)
	Register(dto dto.CreateUser) error
	Login(dto dto.CreateUser) (string, error)
}

type userService struct {
	users map[string]User
}

func NewUserService() UserService {
	users := seedUsers()

	return &userService{users: users}
}

func (s *userService) Get(email string) (User, error) {
	if value, ok := s.users[email]; ok {
		return value, nil
	}

	return User{}, errors.New("user not found")
}

func (s *userService) GetAll() ([]User, error) {
	var users []User

	for _, value := range s.users {
		users = append(users, value)
	}

	return users, nil
}

func (s *userService) Register(dto dto.CreateUser) error {
	if _, ok := s.users[dto.Email]; ok {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	s.users[dto.Email] = User{Email: dto.Email, Password: hashedPassword, Role: RoleUser}

	return nil
}

func (s *userService) Login(dto dto.CreateUser) (string, error) {
	user, err := s.Get(dto.Email)
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte(dto.Password))
	if err != nil {
		return "", err
	}

	jwt, err := token.Generate(user.Email, user.Role)
	if err != nil {
		return "", err
	}

	return jwt, nil
}

func seedUsers() map[string]User {
	users := make(map[string]User)

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	users["admin@mail.com"] = User{Email: "admin@mail.com", Password: hashedPassword, Role: RoleAdmin}

	return users
}
