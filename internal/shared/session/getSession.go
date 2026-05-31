package session

import (
	"fmt"

	"github.com/gorilla/sessions"
	echosession "github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func getSession(c echo.Context) (*sessions.Session, error) {
	sess, err := echosession.Get("session", c)
	if err != nil {
		fmt.Printf("Error retrieving session: %v\n", err)
		return nil, ErrBadRetrieval
	}

	return sess, err
}
