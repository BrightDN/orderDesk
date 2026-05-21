package handlers

import (
	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/shared/parse"
	"github.com/labstack/echo/v4"
)

func (h *Handler) deleteCompanyInvite(c echo.Context) error {
	id, err := parse.Int32(c.Param("id"))
	if err != nil {
		if flashErr := flash.Trigger(c, flash.Error, err.Error()); flashErr != nil {
			return flashErr
		}
		return h.renderInviteListPartial(c)
	}
	if err := h.App.Services.Invitations.Delete(c, id); err != nil {
		if err := flash.Set(c, flash.Error, ErrInternalError.Error()); err != nil {
			return err
		}
		return h.renderInviteListPartial(c)
	}

	if err := flash.Set(c, flash.Pass, "Company invite successfully deleted."); err != nil {
		return err
	}
	return h.renderInviteListPartial(c)
}
