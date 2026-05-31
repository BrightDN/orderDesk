package invites

import (
	"database/sql"
	"errors"
	"time"

	"github.com/labstack/echo/v4"
)

var ErrInvalidInvitationToken = errors.New("invalid invitation token")
var ErrExpiredToken = errors.New("the token has already expired")
var ErrAlreadyUsed = errors.New("this token has already been used")

func (is *InvitationService) GetInvitation(c echo.Context, token string) (Invitation, error) {
	invite, err := is.db.GetInviteByToken(c.Request().Context(), token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return Invitation{}, ErrInvalidInvitationToken
		}
		return Invitation{}, ErrInternalError
	}
	if invite.ExpiresAt.Before(time.Now()) {
		return Invitation{}, ErrExpiredToken
	}
	if invite.UsedAt.Valid {
		return Invitation{}, ErrAlreadyUsed
	}

	var invitation = Invitation{
		Type:    iType(invite.InviteType),
		Email:   invite.Email,
		Company: invite.CompanyName,
		Token:   invite.Token,
	}

	return invitation, nil
}
