package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/labstack/echo/v4"
)

const adminNewCompanyPath = "/admin/companies/invites"

func (h *Handler) sendCompanyInvite(c echo.Context) error {
	err := h.App.Services.Invitations.SendCompany(c)
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return c.Redirect(http.StatusSeeOther, adminNewCompanyPath)
	}

	if logErr := logging.Log_info_and_flash(c, "A company invite has been created and sent", "Company invite created and sent"); logErr != nil {
		return logErr
	}
	return c.Redirect(http.StatusSeeOther, adminNewCompanyPath)
}
