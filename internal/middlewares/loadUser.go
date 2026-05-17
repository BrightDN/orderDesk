package middlewares

import (
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func loadUser(db *database.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, _ := session.Get("session", c)
			userID, ok := sess.Values["userID"].(int32)
			if ok {
				user, err := db.GetUserById(c.Request().Context(), userID)
				if err == nil {
					c.Set("user", user)
				}
			}
			return next(c)
		}
	}
}
