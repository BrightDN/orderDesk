package middlewares

import (
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func LoadUser(cfg *configs.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get("session", c)
			userID, ok := sess.Values["userID"].(int32)
			if ok {
				user, err := cfg.Db.GetUserById(c.Request().Context(), userID)
				if err == nil {
					c.Set("user", user)
				}
			}
			return next(c)
		}
	}
}
