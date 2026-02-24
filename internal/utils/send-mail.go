package utils

import (
	"fmt"
	"log"
	"scheduling-app-back-end/internal/models"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

func SendMsg(m models.MailData) {
	_ = SendMsgErr(m)
}

func SendMsgErr(m models.MailData) error {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	client, err := server.Connect()
	if err != nil {
		log.Println(err)
		return fmt.Errorf("smtp connect: %w", err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, m.Content)

	if err := email.Send(client); err != nil {
		log.Println(err)
		return fmt.Errorf("send mail: %w", err)
	}

	log.Println("Email sent:", m.To)
	return nil
}
