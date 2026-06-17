package authentication

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func (auth *AuthenticationService) Logout(c echo.Context) *errorHandling.AppError {
	if err := session.Clear(c); err != nil {
		return &errorHandling.AppError{
			Action:    "Clearing user session on logout",
			LogError:  err,
			UserError: errors.New("failed to clear session"),
		}
	}
	return nil
}
