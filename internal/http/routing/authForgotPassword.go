package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func (n *Navigation) authForgotPassword(c echo.Context) error {
	return c.Render(http.StatusOK, "/auth/forgot-password", nil)
}

func (n *Navigation) authForgotPasswordRequest(c echo.Context) error {
	if err := flash.Set(c, flash.Pass, "Please contact your administrator to reset your password."); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/auth/login")
}
