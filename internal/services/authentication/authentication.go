package authentication

import (
	"database/sql"

	"github.com/brightDN/orderDesk/internal/database"
)

type AuthenticationService struct {
	queries *database.Queries
	db      *sql.DB
}

func NewAuthService(queries *database.Queries, db *sql.DB) *AuthenticationService {
	return &AuthenticationService{
		queries: queries,
		db:      db,
	}
}
