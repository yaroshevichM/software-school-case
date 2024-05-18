package service

import (
	"fmt"
	"net/smtp"

	"github.com/sirupsen/logrus"
)

type MailService struct {
	host     string
	port     int
	username string
	password string
}

func NewMailService(host string, port int, username, password string) *MailService {
	return &MailService{
		host:     host,
		port:     port,
		username: username,
		password: password,
	}
}

func (s *MailService) SendMail(to, subject, body string) error {
	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	messageString := fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, subject, body)
	msg := []byte(messageString)
	logrus.Printf("Send message to %s", messageString)
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	return smtp.SendMail(addr, auth, s.username, []string{to}, msg)
}
