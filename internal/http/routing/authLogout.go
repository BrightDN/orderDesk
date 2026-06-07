package routing

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (n *Navigation) authLogout(c echo.Context) error {
	if err := n.app.Services.Auth.Logout(c); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
