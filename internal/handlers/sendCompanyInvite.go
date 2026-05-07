package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SendCompanyInvite(c echo.Context) error {
	return c.Redirect(http.StatusSeeOther, "/admin/controlpanel")
}
