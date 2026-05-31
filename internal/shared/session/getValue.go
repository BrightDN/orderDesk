package session

import (
	"github.com/labstack/echo/v4"
)

func GetValue[T any](c echo.Context, key string) (T, bool, error) {
	var zero T

	sess, err := getSession(c)
	if err != nil {
		return zero, false, err
	}

	value, ok := sess.Values[key].(T)
	if !ok {
		return zero, false, nil
	}

	return value, true, nil
}
