package authentication

import "github.com/brightDN/orderDesk/internal/database"

type AuthenticationService struct {
	db *database.Queries
}

func NewAuthService(db *database.Queries) *AuthenticationService {
	return &AuthenticationService{
		db: db,
	}
}
