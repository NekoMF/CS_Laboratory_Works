package user

import (
	"errors"
)

type Database interface {
	Get(key string) (User, error)
	Set(key string, value User) error
	Delete(key string) error
}

type inMemoryDatabase struct {
	data map[string]User
}

func NewInMemoryDatabase() Database {
	return &inMemoryDatabase{data: make(map[string]User)}
}

func (s *inMemoryDatabase) Get(key string) (User, error) {
	if value, ok := s.data[key]; ok {
		return value, nil
	}

	return User{}, errors.New("key not found")
}

func (s *inMemoryDatabase) Set(key string, value User) error {
	s.data[key] = value
	return nil
}

func (s *inMemoryDatabase) Delete(key string) error {

	if _, ok := s.data[key]; ok {
		delete(s.data, key)
		return nil
	}

	return errors.New("key not found")
}
