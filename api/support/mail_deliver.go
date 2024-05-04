package support

import (
	"fmt"
	"go-echo-mail/config"
	"log"
	"net/smtp"
	"strings"
)

type SendMailParam struct {
	From    string
	To      []string
	Subject string
	Body    string
}

// send mail
func SendMail(p *SendMailParam) error {
	conf := config.Mail()
	server := fmt.Sprintf("%s:%s", conf.HostName, conf.Port)
	auth := smtp.CRAMMD5Auth(conf.Username, conf.Password)
	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", strings.Join(p.To, ","), p.Subject, p.Body))

	if err := smtp.SendMail(server, auth, p.From, p.To, msg); err != nil {
		log.Printf("Send mail error: %s", err)
		return err
	}

	return nil
}
