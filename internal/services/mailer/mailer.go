package mailer

import (
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/wneessen/go-mail"
)

type MailerService struct {
	Client *mail.Client
	Email  string
}

func NewMailerService(cfg configs.MailConfig) (*MailerService, error) {
	client, err := NewClient(cfg)
	if err != nil {
		return nil, err
	}

	return &MailerService{
		Client: client,
		Email:  cfg.Email,
	}, nil
}
