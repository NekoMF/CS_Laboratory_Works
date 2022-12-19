package mfa

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

const (
	senderName = "Cryptography"
)

type Mail struct {
	To      []string
	Subject string
	Body    string
}

type IMailService interface {
	Send(mail Mail)
}

type MailService struct {
	from string
	addr string
	auth smtp.Auth
}

func NewMailService() IMailService {

	email := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	addr := smtpHost + ":" + smtpPort
	auth := smtp.PlainAuth("", email, password, smtpHost)

	mailService := &MailService{
		from: email,
		addr: addr,
		auth: auth,
	}

	return mailService
}

func (s *MailService) Send(mail Mail) {
	msg := s.buildMail(mail)

	smtp.SendMail(s.addr, s.auth, s.from, mail.To, msg)
}

func (s *MailService) buildMail(mail Mail) []byte {
	msg := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\r\n"
	msg += fmt.Sprintf("From: %s\r\n", senderName)
	msg += fmt.Sprintf("To: %s\r\n", strings.Join(mail.To, ";"))
	msg += fmt.Sprintf("Subject: %s\r\n", mail.Subject)
	msg += fmt.Sprintf("\r\n%s\r\n", mail.Body)

	return []byte(msg)
}
