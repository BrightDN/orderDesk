package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/services/invites"
	"github.com/labstack/echo/v4"
)

func (h *Handler) reactivateCompanyInvite(c echo.Context) error {
	if err := invites.Reactivate(h.App.Db, c); err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return flashErr
		}
		return h.renderInviteListPartial(c)
	}
	return h.renderInviteListPartial(c)
}

func (h *Handler) renderInviteListPartial(c echo.Context) error {
	invs := invites.GetCompanyInvites(h.App.Db, c)
	return c.Render(http.StatusOK, "partials/inviteList", map[string]any{
		"invites": invs,
	})
}
