package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

func (h *Handler) resendCompanyInvite(c echo.Context) error {
	id, err := parse.Int32(c.Param("id"))
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return c.NoContent(http.StatusNoContent)
	}
	if err := h.App.Services.Invitations.Resend(c, id); err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return c.NoContent(http.StatusNoContent)
	}

	if logErr := logging.Log_info_and_flash_trigger(c, "A company has received a new invitation mail", "Company invite email resent"); logErr != nil {
		return logErr
	}
	return c.NoContent(http.StatusNoContent)
}
