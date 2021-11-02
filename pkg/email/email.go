package email

import (
	"log"
	"net/smtp"
	"strings"
)

type SimpleEmailClient struct {
	Username string
	Password string
	SmtpHost string
	SmtpPort string
	Identity string
}

func checkEmailError(err error) {
	if err != nil {
		log.Printf("checkEmailError: %v", err)
	}
}

func (s SimpleEmailClient) SendMail(fromUser string, toUser []string, subject string, body string, mailType string) error {
	log.Printf("SendMail from: %v toUser: %v, subject: %v, body: %v\n", fromUser, toUser, subject, body)
	msg := []byte(
		"To: " + strings.Join(toUser, ",") +
			"\r\nFrom: Go SendMail Tool<" + fromUser + ">" +
			"\r\nSubject: " + subject +
			"\r\nContent-type: " + mailType + "; charset=UTF-8\r\n\r\n" +
			body)
	auth := smtp.PlainAuth(s.Identity, s.Username, s.Password, s.SmtpHost)
	err := smtp.SendMail(s.SmtpHost+":"+s.SmtpPort, auth, fromUser, toUser, msg)
	if err != nil {
		checkEmailError(err)
		return err
	}
	return nil
}
