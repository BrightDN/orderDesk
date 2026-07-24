package mailer

import (
	"fmt"
	"io"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/wneessen/go-mail"
)

func (ms *MailerService) Send(m Mail) *errorHandling.AppError {
	msg := mail.NewMsg()
	if err := msg.To(m.Receiver); err != nil {
		return &errorHandling.AppError{
			Action:    "Setting the \"to\" field for the mailer",
			LogError:  fmt.Errorf("something went wrong while setting the \"to\" field: %v", err),
			UserError: fmt.Errorf("error: could not send the e-mail. Try again later or contact support"),
		}
	}
	if err := msg.From(ms.email); err != nil {
		return &errorHandling.AppError{
			Action:    "Setting the \"From\" field for the mailer",
			LogError:  fmt.Errorf("something went wrong while setting the \"From\" field: %v", err),
			UserError: fmt.Errorf("error: could not send the e-mail. Try again later or contact support"),
		}
	}
	msg.Subject(m.Subject)
	msg.SetBodyString(mail.TypeTextPlain, m.Body)

	if m.Attachment != nil {
		m.Attachment.Reader.Seek(0, io.SeekStart)
		msg.AttachReadSeeker(m.Attachment.Filename, m.Attachment.Reader)
	}

	if err := ms.client.DialAndSend(msg); err != nil {
		return &errorHandling.AppError{
			Action:    "Dialing and sending email",
			LogError:  fmt.Errorf("something went wrong while dialing and sending the mail: %v", err),
			UserError: fmt.Errorf("error: could not send the e-mail. Try again later or contact support"),
		}
	}
	return nil
}
