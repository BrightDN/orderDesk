package app

import (
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/services"
)

type App struct {
	Services *services.Services
	Db       *database.Queries
	Cfg      configs.Config
	Name     string
}

func New(services *services.Services, db *database.Queries, cfg configs.Config, name string) App {
	return App{
		Services: services,
		Db:       db,
		Cfg:      cfg,
		Name:     name,
	}
}
