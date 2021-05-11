package email

import (
	"fmt"
	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	Subject  string
	SMTPHost string
	SMTP     int
	Sender   string
	Receiver []string
	Code     string
	Call     string
}

type Email struct {
	Subject    string
	Body       string
	Recipients []string
	Sender     string
	Code       string
	SMTPHost   string
	SMTP       int
	Call       string
}

func New(config *EmailConfig) *Email {
	return &Email{
		Subject:    config.Subject,
		Body:       "",
		Recipients: config.Receiver,
		Sender:     config.Sender,
		Code:       config.Code,
		SMTPHost:   config.SMTPHost,
		SMTP:       config.SMTP,
		Call:       config.Call,
	}
}

func (e *Email) Send(subject, body string) error {
	e.Body = body
	message := gomail.NewMessage()
	message.SetHeader("From", e.Sender)
	message.SetHeader("To", e.Recipients...)
	if len(subject) > 0 {
		message.SetHeader("Subject", subject)
	} else {
		message.SetHeader("Subject", e.Subject)
	}
	message.SetBody("text/html", fmt.Sprintf(e.Call, e.Body))
	err := gomail.NewDialer(e.SMTPHost, e.SMTP, e.Sender, e.Code).DialAndSend(message)
	if err != nil {
		return err
	}
	return nil
}
