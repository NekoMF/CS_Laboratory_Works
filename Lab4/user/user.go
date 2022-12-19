package user

import "crypto/rsa"

type User struct {
	Email      string
	Password   []byte
	PrivateKey *rsa.PrivateKey
}
