package invites

import (
	"database/sql"
	"errors"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

var ErrInvalidInvitation = errors.New("invalid invitation")
var ErrInternalServerError = errors.New("internal server error")

func (inv *InvitationService) ValidateInvitation(c echo.Context, token, email string) (*database.Invite, error) {
	invite, err := inv.db.ValidateInvite(c.Request().Context(), database.ValidateInviteParams{
		Token: token,
		Email: email,
	})

	if err == sql.ErrNoRows {
		return nil, ErrInvalidInvitation
	}
	if err != nil {
		return nil, ErrInternalServerError
	}
	return &invite, nil
}
