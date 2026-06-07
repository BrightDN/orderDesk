package authentication

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

var ErrSessionErr = errors.New("error: Failed clearing session")

func (auth *AuthenticationService) Logout(c echo.Context) error {
	if err := session.Clear(c); err != nil {
		flashErr := flash.Set(c, flash.Error, ErrSessionErr.Error())
		if flashErr != nil {
			return flashErr
		}
		return err
	}
	return nil
}
