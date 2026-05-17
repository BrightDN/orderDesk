package handlers

import (
	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/services/invites"
	"github.com/labstack/echo/v4"
)

func (h *Handler) deleteCompanyInvite(c echo.Context) error {
	if err := invites.Delete(h.App.Db, c); err != nil {
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
