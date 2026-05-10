package mailer

import (
	"fmt"
	"strings"

	"github.com/wneessen/go-mail"
)

func NewClient(provider string, port int, username, password string) (*mail.Client, error) {
	if len(strings.TrimSpace(provider)) == 0 {
		return nil, fmt.Errorf("Mail provider can not be empty")
	}
	if len(strings.TrimSpace(username)) == 0 {
		return nil, fmt.Errorf("username can not be empty")
	}
	if len(strings.TrimSpace(password)) == 0 {
		return nil, fmt.Errorf("password can not be empty")
	}

	if port < 0 {
		port = 587
	}
	c, err := mail.NewClient(
		provider,
		mail.WithPort(port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(username),
		mail.WithPassword(password),
	)

	if err != nil {
		return nil, fmt.Errorf("Something went wrong initializing the mailing provider: %v", err)
	}

	return c, nil
}
