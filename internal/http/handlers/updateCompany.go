package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

func (h *Handler) updateCompany(c echo.Context) error {
	id, err := parse.Int32(c.Param("id"))
	if err != nil {
		return c.Redirect(http.StatusSeeOther, "/admin/companies/overview")
	}

	if err := h.App.Services.Companies.Update(c, id); err != nil {
		return err
	}

	return c.Redirect(http.StatusSeeOther, "/admin/companies/details/"+c.Param("id"))
}
