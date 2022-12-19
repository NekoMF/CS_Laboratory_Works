package mfa

import (
	"errors"
	"math/rand"
	"strconv"
)

type IOtpService interface {
	Generate(email string) (string, error)
	Verify(email string, otp string) error
}

type OtpService struct {
	otps map[string]string
}

func NewOtpService() IOtpService {

	return &OtpService{
		otps: make(map[string]string),
	}
}

func (s *OtpService) Generate(email string) (string, error) {
	num := 100000 + rand.Intn(800000)
	otp := strconv.Itoa(num)

	s.otps[email] = otp

	return otp, nil
}

func (s *OtpService) Verify(email string, otp string) error {
	if s.otps[email] != otp {
		return errors.New("otp is not valid")
	}

	delete(s.otps, email)

	return nil
}
