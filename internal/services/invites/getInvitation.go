package invites

import (
	"database/sql"
	"errors"
	"time"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

var ErrInvalidInvitation = errors.New("invalid invitation token")
var ErrExpiredToken = errors.New("the token has already expired")
var ErrAlreadyUsed = errors.New("this token has already been used")

func (is *InvitationService) GetInvitation(c echo.Context, token string) (Invitation, error) {
	invite, err := is.db.GetInviteByToken(c.Request().Context(), token)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			if flashErr := flash.Set(c, flash.Error, ErrInvalidInvitation.Error()); flashErr != nil {
				return Invitation{}, flashErr
			}
			return Invitation{}, ErrInvalidInvitation
		}
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return Invitation{}, flashErr
		}
		return Invitation{}, ErrInternalError
	}
	if invite.ExpiresAt.Before(time.Now()) {
		if flashErr := flash.Set(c, flash.Error, ErrExpiredToken.Error()); flashErr != nil {
			return Invitation{}, flashErr
		}
		return Invitation{}, ErrExpiredToken
	}
	if invite.UsedAt.Valid {
		if flashErr := flash.Set(c, flash.Error, ErrAlreadyUsed.Error()); flashErr != nil {
			return Invitation{}, flashErr
		}
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
