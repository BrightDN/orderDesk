package invites

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

var ErrTokenCreation = errors.New("failed to generate a token")

func generateToken(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("%w: %v", ErrTokenCreation, err)
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
