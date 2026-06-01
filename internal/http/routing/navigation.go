package routing

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/app"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/middlewares"
	"github.com/labstack/echo/v4"
)

type Navigation struct {
	app *app.App
	db  *database.Queries
}

func NewNav(db *database.Queries, app *app.App) *Navigation {
	return &Navigation{
		db:  db,
		app: app,
	}
}

func (n *Navigation) Register(e *echo.Echo) {
	withEmployee := []echo.MiddlewareFunc{
		middlewares.RequireAuth(),
		middlewares.LoadEmployee(n.db),
	}

	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusSeeOther, "/app/neworder")
	})

	e.GET("/dashboard/suppliers", func(c echo.Context) error {
		return c.Render(http.StatusOK, "suppliers", map[string]any{
			"Page": "suppliers",
		})
	}, withEmployee...)

	e.GET("/dashboard/history", func(c echo.Context) error {
		return c.Render(http.StatusOK, "orderHistory", map[string]any{
			"Page": "history",
		})
	}, withEmployee...)

	e.GET("/dashboard/settings", func(c echo.Context) error {
		return c.Render(http.StatusOK, "companySettings", map[string]any{
			"Page": "company settings",
		})
	}, withEmployee...)

	e.GET("/support/contact", func(c echo.Context) error {
		return c.Render(http.StatusOK, "contactPage", nil)
	}, withEmployee...)

	e.GET("/settings/companies", func(c echo.Context) error {
		return c.Render(http.StatusOK, "companies", nil)
	}, withEmployee...)

	e.GET("/settings/user", func(c echo.Context) error {
		return c.Render(http.StatusOK, "userSettings", nil)
	}, withEmployee...)

	// Business
	e.GET("/app/neworder", n.appNewOrder, withEmployee...)

	// Site admin
	e.GET("/admin/companies/invites", n.adminCompanyInvite, withEmployee...)
	e.GET("/admin/companies/overview", n.adminCompanyOverview, withEmployee...)
	e.GET("/admin/companies/details/:id", n.adminCompanyDetails, withEmployee...)

	// Authentication
	e.GET("/auth/signup/:token", n.authSignUp)
	e.GET("/auth/login", n.authLogin)
}
