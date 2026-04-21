package configs

import "github.com/brightDN/orderDesk/internal/database"

type Config struct {
	Db       *database.Queries
	Platform string
}
