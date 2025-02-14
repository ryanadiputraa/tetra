package mail

import (
	"context"
	"fmt"
	"net/smtp"
)

type SMTPMail interface {
	SendMail(ctx context.Context, to, subject, body string) error
}

type mail struct {
	host   string
	port   string
	sender string
	pass   string
}

func NewSMTPMail(sender, pass string) SMTPMail {
	return &mail{
		host:   "smtp.gmail.com",
		port:   "587",
		sender: sender,
		pass:   pass,
	}
}

func (m *mail) SendMail(ctx context.Context, to, subject, body string) (err error) {
	auth := smtp.PlainAuth("", m.sender, m.pass, m.host)
	addr := fmt.Sprintf("%v:%v", m.host, m.port)
	msg := fmt.Sprintf("Subject: %s\nMIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n%s", subject, body)
	err = smtp.SendMail(addr, auth, m.sender, []string{to}, []byte(msg))
	return
}
