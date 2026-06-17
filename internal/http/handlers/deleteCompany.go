package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

func (h *Handler) deleteCompany(c echo.Context) error {
	id, err := parse.Int32(c.Param("id"))
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return h.renderCompanyListPartial(c)
	}
	if err := h.App.Services.Companies.Delete(c, id); err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return h.renderCompanyListPartial(c)
	}

	if err := logging.Log_info_and_flash(c, "A company has been deleted", "Company successfully removed."); err != nil {
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
