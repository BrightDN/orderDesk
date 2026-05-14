package flash

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Set(c echo.Context, messageType MessageType, message string) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	sess.AddFlash(string(messageType)+separator+message, feedbackKey)

	return sess.Save(c.Request(), c.Response())
}
