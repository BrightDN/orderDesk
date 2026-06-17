package invites

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
)

func (is *InvitationService) generateToken(length int) (string, *errorHandling.AppError) {
	bytes := make([]byte, length)

	if _, err := rand.Read(bytes); err != nil {
		return "", &errorHandling.AppError{
			Action:    "Generating invitation token",
			LogError:  fmt.Errorf("Failed to read random bytes: %v", err),
			UserError: errors.New("failed to generate token"),
		}
	}

	return base64.RawURLEncoding.EncodeToString(bytes), nil
}
