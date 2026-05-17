package mailer

import (
	"fmt"
	"strings"

	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/wneessen/go-mail"
)

func NewClient(cfg configs.MailConfig) (*mail.Client, error) {
	if len(strings.TrimSpace(cfg.Provider)) == 0 {
		return nil, fmt.Errorf("Mail provider can not be empty")
	}
	if len(strings.TrimSpace(cfg.Username)) == 0 {
		return nil, fmt.Errorf("username can not be empty")
	}
	if len(strings.TrimSpace(cfg.Password)) == 0 {
		return nil, fmt.Errorf("password can not be empty")
	}

	if cfg.Port < 0 {
		cfg.Port = 587
	}
	c, err := mail.NewClient(
		cfg.Provider,
		mail.WithPort(cfg.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.Username),
		mail.WithPassword(cfg.Password),
	)

	if err != nil {
		return nil, fmt.Errorf("Something went wrong initializing the mailing provider: %v", err)
	}

	return c, nil
}
