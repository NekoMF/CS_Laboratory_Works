package message

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
)

type MessageService interface {
	SignMessage(privateKey *rsa.PrivateKey, msg []byte) ([]byte, []byte, error)
	VerifyMessage(pubKey *rsa.PublicKey, message, signature []byte) error
}

type messageService struct{}

func NewMessageService() MessageService {
	return &messageService{}
}

func (s *messageService) SignMessage(privateKey *rsa.PrivateKey, msg []byte) ([]byte, []byte, error) {

	// Before signing, we need to hash our message
	// The hash is what we actually sign
	msgHash := sha256.New()
	_, err := msgHash.Write(msg)
	if err != nil {
		return nil, nil, err
	}
	msgHashSum := msgHash.Sum(nil)

	// In order to generate the signature, we provide a random number generator,
	// our private key, the hashing algorithm that we used, and the hash sum
	// of our message
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		return nil, nil, err
	}

	return signature, msgHashSum, nil
}

func (s *messageService) VerifyMessage(publicKey *rsa.PublicKey, signature []byte, msgHashSum []byte) error {

	// To verify the signature, we provide the public key, the hashing algorithm
	// the hash sum of our message and the signature we generated previously
	// there is an optional "options" parameter which can omit for now
	return rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
}
