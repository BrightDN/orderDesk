package handlers

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/labstack/echo/v4"
)

func (h *Handler) reactivateCompanyInvite(c echo.Context) error {
	if err := h.App.Services.Invitations.Reactivate(c); err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return flashErr
		}
		return h.renderInviteListPartial(c)
	}
	return h.renderInviteListPartial(c)
}

func (h *Handler) renderInviteListPartial(c echo.Context) error {
	invs := h.App.Services.Invitations.GetCompanyInvites(c)
	return c.Render(http.StatusOK, "partials/inviteList", map[string]any{
		"invites": invs,
		"csrf":    c.Get("csrf"),
	})
}
