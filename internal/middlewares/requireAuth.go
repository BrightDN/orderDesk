package middlewares

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

func RequireAuth() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			_, ok, err := session.GetValue[int32](c, session.UserIDKey)
			if err != nil {
				return err
			}
			if !ok {
				return c.Redirect(http.StatusSeeOther, "/auth/login")
			}
			return next(c)
		}
	}
}
