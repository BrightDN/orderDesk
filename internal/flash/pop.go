package flash

import (
	"strings"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Pop(c echo.Context) (*Flash, error) {
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

	return &Flash{
		Type:    MessageType(messageType),
		Message: message,
	}, sess.Save(c.Request(), c.Response())
}
