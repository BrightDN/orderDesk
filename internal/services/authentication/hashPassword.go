package authentication

import (
	"errors"
	"fmt"

	"github.com/alexedwards/argon2id"
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
)

func (auth *AuthenticationService) hashPassword(password string) (string, *errorHandling.AppError) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", &errorHandling.AppError{
			Action:    "Hashing password",
			LogError:  fmt.Errorf("Failed to hash password: %v", err),
			UserError: errors.New("failed to hash password"),
		}
	}
	return hash, nil
}
