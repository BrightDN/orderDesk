package handlers

import (
	"github.com/brightDN/orderDesk/internal/shared/errorHandling"
	"github.com/brightDN/orderDesk/internal/shared/logging"
	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

func (h *Handler) deleteCompanyInvite(c echo.Context) error {
	id, err := parse.Int32(c.Param("id"))
	if err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return h.renderInviteListPartial(c)
	}
	if err := h.App.Services.Invitations.Delete(c, id); err != nil {
		if logErr := errorHandling.Log_and_flash(c, *err); logErr != nil {
			return logErr
		}
		return h.renderInviteListPartial(c)
	}

	if err := logging.Log_info_and_flash(c, "Deleted company invitation", "Company invite successfully deleted."); err != nil {
		return err
	}
	return h.renderInviteListPartial(c)
}
