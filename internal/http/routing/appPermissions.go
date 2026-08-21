package routing

import (
	"github.com/brightDN/orderDesk/internal/services/permissions"
	"github.com/labstack/echo/v4"
)

func appPermissions(c echo.Context) permissions.Permissions {
	perms, _ := c.Get(permissions.ContextKey).(permissions.Permissions)
	return perms
}
