package app

import (
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/wneessen/go-mail"
)

type App struct {
	Db     *database.Queries
	Cfg    configs.Config
	Mailer *mail.Client
}

func New(db *database.Queries, cfg configs.Config, mailer *mail.Client) App {
	return App{
		Db:     db,
		Cfg:    cfg,
		Mailer: mailer,
	}
}
