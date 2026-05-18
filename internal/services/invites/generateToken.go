package invites

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func (is *InvitationService) generateToken(length int) (string, error) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("%w: %v", ErrTokenCreation, err)
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
