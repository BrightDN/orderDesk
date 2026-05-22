package authentication

import (
	"fmt"

	"github.com/alexedwards/argon2id"
)

func (auth *AuthenticationService) hashPassword(password string) (string, error) {
	hash, err := argon2id.CreateHash(password, argon2id.DefaultParams)
	if err != nil {
		return "", fmt.Errorf("something went wrong: %v", err)
	}
	return hash, nil
}
