package handlers

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/brightDN/orderDesk/internal/flash"
	"github.com/brightDN/orderDesk/internal/invites"
	"github.com/labstack/echo/v4"
)

func (h *Handler) DeleteCompanyInvite(c echo.Context) error {
	id := c.Param("id")
	invs := invites.GetCompanyInvites(h.App.Db, c, h.App.Name)
	if len(strings.TrimSpace(id)) == 0 {
		if flashErr := flash.Set(c, "error", ErrUnexpectedValue.Error()); flashErr != nil {
			return flashErr
		}
		return c.Render(http.StatusOK, "partials/inviteList", map[string]any{
			"invites": invs,
		})
	}

	nid, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		if flashErr := flash.Set(c, "error", ErrUnexpectedValue.Error()); flashErr != nil {
			return flashErr
		}
		return c.Render(http.StatusOK, "partials/inviteList", map[string]any{
			"invites": invs,
		})
	}

	if err := invites.Delete(h.App.Db, c, int32(nid)); err != nil {
		if err := flash.Set(c, "error", ErrInternalError.Error()); err != nil {
			return err
		}
		return c.Render(http.StatusOK, "partials/inviteList", map[string]any{
			"invites": invs,
		})
	}

	if err := flash.Set(c, "pass", "Company invite succesfully deleted."); err != nil {
		return err
	}
	return c.Render(http.StatusOK, "partials/inviteList", map[string]any{
		"invites": invs,
	})
}
