package mailx

import (
	"fmt"
	"net/smtp"
)

type SMTPMail struct {
	user     string //example@example.com login smtp server user
	password string //login smtp server password
	host     string //smtp.163.com/smtp.qq.com/...
	port     uint   //25/587/...
}

func NewSMTPMail(user, password, host string, port uint) SMTPMail {
	return SMTPMail{
		user:     user,
		password: password,
		host:     host,
		port:     port,
	}
}

type mailContentType string

const (
	ContentTypeTextPlain mailContentType = "Content-Type: text/plain; charset=UTF-8"
	ContentTypeTextHtml  mailContentType = "Content-Type: text/html; charset=UTF-8"
)

func (mail SMTPMail) send(subject, body string, mailType mailContentType, receivers []string) error {
	host := mail.host
	port := mail.port
	user := mail.user
	password := mail.password
	addr := fmt.Sprintf("%s:%d", host, port)

	var contentType = string(mailType)
	msg := []byte("From: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	auth := smtp.PlainAuth("", user, password, host)
	err := smtp.SendMail(addr, auth, user, receivers, msg)
	return err
}

func (mail SMTPMail) Send(receivers []string, subject, body string, mailContentType mailContentType) error {
	err := mail.send(subject, body, mailContentType, receivers)
	return err
}
