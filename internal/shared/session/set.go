package session

import (
	"reflect"

	"github.com/labstack/echo/v4"
)

func SetValues(c echo.Context, data SessionData) error {
	sess, err := getSession(c)
	if err != nil {
		return err
	}

	if sess.Values == nil {
		sess.Values = make(map[any]any)
	}

	val := reflect.ValueOf(data)
	typ := reflect.TypeFor[SessionData]()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		sess.Values[field.Name] = val.Field(i).Interface()
	}

	return sess.Save(c.Request(), c.Response())
}
