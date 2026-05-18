package mailer

import (
	"errors"
	"fmt"

	"github.com/wneessen/go-mail"
)

var ErrMailSending = errors.New("failed to send mail")

func (ms *MailerService) Send(m Mail) error {
	msg := mail.NewMsg()
	if err := msg.To(m.Receiver); err != nil {
		return fmt.Errorf("%w: %v", ErrMailSending, err)
	}
	if err := msg.From(ms.Email); err != nil {
		return fmt.Errorf("%w: %v", ErrMailSending, err)
	}
	msg.Subject(m.Subject)
	msg.SetBodyString(mail.TypeTextPlain, m.Body)

	for _, attachment := range m.Attachments {
		msg.AttachFile(attachment)
	}
	if err := ms.Client.DialAndSend(msg); err != nil {
		return fmt.Errorf("%w: %v", ErrMailSending, err)
	}
	return nil
}
