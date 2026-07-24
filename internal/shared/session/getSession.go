package session

import (
	"fmt"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/gorilla/sessions"
	echosession "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func getSession(c echo.Context) (*sessions.Session, *errorHandling.AppError) {
	sess, err := echosession.Get("session", c)
	if err != nil {
		appErr := errorHandling.AppError{
			Action:    "getting session",
			LogError:  fmt.Errorf("error retrieving session: %w", err),
			UserError: fmt.Errorf("error retrieving session"),
		}
		return nil, &appErr
	}
	return sess, nil
}
