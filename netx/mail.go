package netx

import (
	"fmt"
	"net/smtp"
	"strings"
)

type Mailer struct {
  username string
  password string
  host 	   string
  port     string
  nickname string
}

func (m *Mailer) SendMail(to []string, subject, body string) bool {
	auth := smtp.PlainAuth("", m.username, m.password, m.host)
	content_type := "Content-Type: text/plain; charset=UTF-8"
	msg := []byte("To: " + strings.Join(to, ",") + "\r\nFrom: " + m.nickname +
		"<" + m.username + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	err := smtp.SendMail(m.host+":"+m.port, auth, m.username, to, msg)
	if err != nil {
		fmt.Printf("send mail error: %v", err)
		return false
	}
	return true
}
