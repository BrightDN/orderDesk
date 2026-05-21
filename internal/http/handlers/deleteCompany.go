package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

func (h *Handler) deleteCompany(c echo.Context) error {
	id, err := parse.Int32(c.Param("id"))
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return h.renderCompanyListPartial(c)
	}
	if err := h.App.Services.Companies.Delete(c, id); err != nil {
		if flashErr := flash.Set(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return h.renderCompanyListPartial(c)
	}

	if err := flash.Set(c, flash.Pass, "Company successfully removed."); err != nil {
		return err
	}
	return h.renderCompanyListPartial(c)
}

func (h *Handler) renderCompanyListPartial(c echo.Context) error {
	companyList, err := h.App.Services.Companies.GetCompanies(c)
	if err != nil {
		return err
	}
	return c.Render(http.StatusOK, "partials/companyList", map[string]any{
		"companies": companyList,
	})
}
