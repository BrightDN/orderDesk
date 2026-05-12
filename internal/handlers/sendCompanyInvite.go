package handlers

import (
	"errors"
	"net/http"

	"github.com/brightDN/orderDesk/internal/companies"
	"github.com/brightDN/orderDesk/internal/invites"
	"github.com/labstack/echo/v4"
)

var ErrDBDataRetrieval = errors.New("Something went wrong retrieving data")

func (h *Handler) SendCompanyInvite(c echo.Context) error {

	err := invites.SendCompany(h.App.Db, c, h.App.Name, h.App.Cfg.MailAccount, h.App.Mailer)
	invs := invites.GetCompanyInvites(h.App.Db, c, h.App.Name)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, companies.ErrDuplicateEmail) ||
			errors.Is(err, companies.ErrDuplicateName) ||
			errors.Is(err, invites.ErrMaxAttempts) {
			statusCode = http.StatusUnprocessableEntity
		}
		return c.Render(statusCode, "inviteCompany", map[string]any{
			"feedback": map[string]string{
				"message": err.Error(),
				"type":    "error",
			},
			"invites": invs,
		})
	}
	return c.Render(http.StatusOK, "inviteCompany", map[string]any{
		"feedback": map[string]string{
			"message": "Company invite created and sent.",
			"type":    "pass",
		},
		"invites": invs,
	})
}
