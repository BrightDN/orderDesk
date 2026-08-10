package authentication

import (
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (auth *AuthenticationService) Logout(c echo.Context) error {
	if err := session.Clear(c); err != nil {
		return err
	}
	return nil
}
