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
	e.GET(Neworder, n.appNewOrder, withEmployee...).Name = "app.new-order"
	e.GET(Csuppliers, n.appSuppliers, withEmployee...).Name = "app.suppliers"
	e.GET("/app/history", n.appOrderHistory, withEmployee...).Name = "app.history"
	e.GET("/app/settings/company", n.appCompanySettings, withEmployee...).Name = "app.settings.company"
	e.GET("/app/settings/user", n.appUserSettings, withEmployee...).Name = "app.settings.user"
	e.GET("/app/suppliers/get/:supplier-name", n.appGetSupplier, withEmployee...).Name = "app.suppliers.get"

	// Site admin
	e.GET("/admin/companies/invites", n.adminCompanyInvite, withOwner...).Name = "admin.companies.invites"
	e.GET("/admin/companies/overview", n.adminCompanyOverview, withOwner...).Name = "admin.companies.overview"
	e.GET("/admin/companies/details/:id", n.adminCompanyDetails, withOwner...).Name = "admin.companies.details"

	// Authentication
	e.GET(Login, n.authLogin).Name = "auth.login" // TODO: redirect to dashboard if already logged in
	e.GET(Logout, n.authLogout, middlewares.RequireAuth()).Name = "auth.logout"
	e.GET(Signup, n.authSignUp).Name = "auth.signup"
	e.GET("/auth/forgot-password", n.authForgotPassword).Name = "auth.forgot-password"
	e.POST("/auth/forgot-password", n.authForgotPasswordRequest).Name = "auth.forgot-password.request"
	e.GET("/auth/select-company", n.authSelectCompany, middlewares.RequireAuth()).Name = "auth.select-company"

	// support, TODO: add FAQ, TOS&PP, landing page
	e.GET("/", n.authLogin).Name = "root" // TODO: Create a landing page and redirect to it instead of login
	e.GET("/support/contact", n.supportContact).Name = "support.contact"
}
