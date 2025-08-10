package massager

import (
	"log"

	"github.com/Maryam-nokohan/GoldCurrentPrice/model"
	gomail "gopkg.in/mail.v2"
)

type emailMessenger struct {
	From string
	Key  string
}

func NewMail(from, key string) Massenger {
	if key == "" || from == "" {
		log.Fatal("key or from is empty")
	}

	return &emailMessenger{
		From: from,
		Key:  key,
	}
}

func (e *emailMessenger) Send(to string, message *model.Message) error {
	msg := gomail.NewMessage()

	msg.SetHeader("From", e.From)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", message.Header)
	msg.SetBody("text/plain", message.Body)

	dialer := gomail.NewDialer("smtp.gmail.com", 587, e.From, e.Key)

	if err := dialer.DialAndSend(msg); err != nil {
		log.Println(err)
		return ErrInSending
	} else {
		log.Println("Email sent successfully!")
	}
	return nil
}
