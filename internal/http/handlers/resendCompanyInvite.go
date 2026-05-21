package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

func (h *Handler) resendCompanyInvite(c echo.Context) error {
	id, err := parse.Int32(c.Param("id"))
	if err != nil {
		if flashErr := flash.Trigger(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return c.NoContent(http.StatusNoContent)
	}
	if err := h.App.Services.Invitations.Resend(c, h.App.Name, id); err != nil {
		if flashErr := flash.Trigger(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return c.NoContent(http.StatusNoContent)
	}

	if err := flash.Trigger(c, flash.Pass, "Company invite email resent."); err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
