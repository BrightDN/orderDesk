package mailer

import (
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/wneessen/go-mail"
)

type MailerService struct {
	client *mail.Client
	email  string
}

func NewMailerService(cfg configs.MailConfig) (*MailerService, error) {
	client, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &MailerService{
		client: client,
		email:  cfg.Email,
	}, nil
}
