package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) authSignUp(c echo.Context) error {

	invitation, err := n.app.Services.Invitations.GetInvitation(c, c.Param("token"))
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return c.Redirect(http.StatusSeeOther, "/auth/login")
	}

	return c.Render(http.StatusOK, "authSignUp", map[string]any{
		"invitation": invitation,
	})
}
