package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/invites"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Reactivate(c echo.Context) error {
	id := c.Param("id")
	if len(strings.TrimSpace(id)) == 0 {
		if flashErr := flash.Set(c, flash.Error, ErrUnexpectedValue.Error()); flashErr != nil {
			return flashErr
		}
		invs := invites.GetCompanyInvites(h.App.Db, c, h.App.Name)
		return c.Render(http.StatusOK, "partials/inviteList", map[string]any{
			"invites": invs,
		})
	}

	nid, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrUnexpectedValue.Error()); flashErr != nil {
			return flashErr
		}
		invs := invites.GetCompanyInvites(h.App.Db, c, h.App.Name)
		return c.Render(http.StatusOK, "partials/inviteList", map[string]any{
			"invites": invs,
		})
	}

	if err := invites.Reactivate(h.App.Db, c, int32(nid)); err != nil {
		if flashErr := flash.Set(c, flash.Error, ErrInternalError.Error()); flashErr != nil {
			return flashErr
		}
		invs := invites.GetCompanyInvites(h.App.Db, c, h.App.Name)
		return c.Render(http.StatusOK, "partials/inviteList", map[string]any{
			"invites": invs,
		})
	}

	invs := invites.GetCompanyInvites(h.App.Db, c, h.App.Name)
	return c.Render(http.StatusOK, "partials/inviteList", map[string]any{
		"invites": invs,
	})
}
