package routing

import (
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

const (
	Login  = "/auth/login"
	Logout = "/auth/logout"
	Signup = "auth/signup/:token"

	Neworder   = "/app/neworder"
	Csuppliers = "/app/suppliers"
)

func (n *Navigation) Register(e *echo.Echo) {
	withEmployee := []echo.MiddlewareFunc{
		middlewares.RequireAuth(),
		middlewares.LoadEmployee(n.db),
		// middlewares.LoadPermissions(n.db), // TODO: implement this middleware
	}

	withOwner := []echo.MiddlewareFunc{
		middlewares.RequireAuth(),
		middlewares.RequireOwner(n.db),
	}

	// Business
	e.GET(Neworder, n.appNewOrder, withEmployee...)
	e.GET(Csuppliers, n.appSuppliers, withEmployee...)
	e.GET("/app/history", n.appOrderHistory, withEmployee...)
	e.GET("/app/settings/company", n.appCompanySettings, withEmployee...)
	e.GET("/app/settings/user", n.appUserSettings, withEmployee...)
	e.GET("/app/suppliers/get/:supplier-name", n.appGetSupplier, withEmployee...)

	// Site admin
	e.GET("/admin/companies/invites", n.adminCompanyInvite, withOwner...)
	e.GET("/admin/companies/overview", n.adminCompanyOverview, withOwner...)
	e.GET("/admin/companies/details/:id", n.adminCompanyDetails, withOwner...)

	// Authentication
	e.GET(Login, n.authLogin) // TODO: redirect to dashboard if already logged in
	e.GET(Logout, n.authLogout, middlewares.RequireAuth())
	e.GET(Signup, n.authSignUp)
	e.GET("/auth/forgot-password", n.authForgotPassword)
	e.POST("/auth/forgot-password", n.authForgotPasswordRequest)
	e.GET("/auth/select-company", n.authSelectCompany, middlewares.RequireAuth())

	// support, TODO: add FAQ, TOS&PP, landing page
	e.GET("/", n.authLogin) // TODO: Create a landing page and redirect to it instead of login
	e.GET("/support/contact", n.supportContact)
}
