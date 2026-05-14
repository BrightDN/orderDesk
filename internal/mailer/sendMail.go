package mailer

import (
	"errors"
	"fmt"

	"github.com/wneessen/go-mail"
)

var ErrMailSending = errors.New("failed to send mail")

func SendMail(m Mail, client *mail.Client) error {
	msg := mail.NewMsg()
	if err := msg.To(m.Receiver); err != nil {
		return fmt.Errorf("%w: %v", ErrMailSending, err)
	}
	if err := msg.From(m.Sender); err != nil {
		return fmt.Errorf("%w: %v", ErrMailSending, err)
	}
	msg.Subject(m.Subject)
	msg.SetBodyString(mail.TypeTextPlain, m.Body)

	if err := client.DialAndSend(msg); err != nil {
		return fmt.Errorf("%w: %v", ErrMailSending, err)
	}
	return nil
}
