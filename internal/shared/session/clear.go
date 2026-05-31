package session

import (
	"github.com/labstack/echo/v4"
)

func Clear(c echo.Context) error {
	sess, err := getSession(c)
	if err != nil {
		return err
	}
	for k := range sess.Values {
		delete(sess.Values, k)
	}
	return sess.Save(c.Request(), c.Response())
}
