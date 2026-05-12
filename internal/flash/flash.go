package flash

import (
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const feedbackKey = "feedback"
const separator = "\x00"

func Set(c echo.Context, messageType string, message string) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}

	sess.AddFlash(messageType+separator+message, feedbackKey)

	return sess.Save(c.Request(), c.Response())
}

func Pop(c echo.Context) (map[string]string, error) {
	sess, err := session.Get("session", c)
	if err != nil {
		return nil, err
	}

	flashes := sess.Flashes(feedbackKey)
	if len(flashes) == 0 {
		return nil, nil
	}

	value, _ := flashes[0].(string)
	messageType, message, _ := strings.Cut(value, separator)

	return map[string]string{
		"type":    messageType,
		"message": message,
	}, sess.Save(c.Request(), c.Response())
}
