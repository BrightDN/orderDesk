package authentication

import (
	"fmt"

	"github.com/alexedwards/argon2id"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
)

func (auth *AuthenticationService) comparePasswordHash(password string, hash string) (bool, *errorHandling.AppError) {
	isSame, err := argon2id.ComparePasswordAndHash(password, hash)
	if err != nil {
		return false, &errorHandling.AppError{
			Action:   "Comparing password hash",
			LogError: fmt.Errorf("Failed to compare passwords: %v", err),
		}
	}
	if !isSame {
		return false, nil
	}
	return true, nil
}
