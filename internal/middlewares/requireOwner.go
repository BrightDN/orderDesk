package middlewares

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/shared/session"
	"github.com/labstack/echo/v4"
)

const ownerRole = "site_admin"

func RequireOwner(_ *database.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			role, ok, err := session.GetValue[string](c, session.RoleNameKey)
			if err != nil {
				return err
			}
			if !ok {
				return c.Redirect(http.StatusSeeOther, "/auth/login")
			}
			if role != ownerRole {
				return echo.NewHTTPError(http.StatusForbidden)
			}
			return next(c)
		}
	}
}
