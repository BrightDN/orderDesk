package routing

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (n *Navigation) authLogin(c echo.Context) error {
	return c.Render(http.StatusOK, "authLogin", map[string]any{
		"identity": n.identity,
	})
}
