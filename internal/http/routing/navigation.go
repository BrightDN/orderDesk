package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/app"
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/labstack/echo/v4"
)

type Navigation struct {
	app      *app.App
	db       *database.Queries
	identity *configs.IdentityConfig
}

func NewNav(db *database.Queries, app *app.App, identity *configs.IdentityConfig) *Navigation {
	return &Navigation{
		db:       db,
		app:      app,
		identity: identity,
	}
}

func (n *Navigation) Register(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "newOrder", map[string]any{
			"Page": "new order",
		})
	})

	e.GET("/dashboard/suppliers", func(c echo.Context) error {
		return c.Render(http.StatusOK, "suppliers", map[string]any{
			"Page": "suppliers",
		})
	})

	e.GET("/dashboard/neworder", func(c echo.Context) error {
		return c.Render(http.StatusOK, "newOrder", map[string]any{
			"Page": "new order",
		})
	})

	e.GET("/dashboard/history", func(c echo.Context) error {
		return c.Render(http.StatusOK, "orderHistory", map[string]any{
			"Page": "history",
		})
	})

	e.GET("/dashboard/settings", func(c echo.Context) error {
		return c.Render(http.StatusOK, "companySettings", map[string]any{
			"Page": "company settings",
		})
	})

	e.GET("/support/contact", func(c echo.Context) error {
		return c.Render(http.StatusOK, "contactPage", nil)
	})

	e.GET("/auth/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login", nil)
	})

	e.GET("/auth/signup", func(c echo.Context) error {
		return c.Render(http.StatusOK, "signup", nil)
	})

	e.GET("/settings/companies", func(c echo.Context) error {
		return c.Render(http.StatusOK, "companies", nil)
	})

	e.GET("/settings/user", func(c echo.Context) error {
		return c.Render(http.StatusOK, "userSettings", nil)
	})

	e.GET("/admin/companies/invites", n.adminCompanyInvite)
	e.GET("/admin/companies/overview", n.adminCompanyOverview)
	e.GET("/admin/companies/details/:id", n.adminCompanyDetails)

	e.GET("/auth/company/:token", n.authSignUp)
}
