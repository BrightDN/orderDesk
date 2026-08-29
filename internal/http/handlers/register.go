package handlers

import (
	"github.com/brightDN/orderDesk/internal/middlewares"
	"github.com/labstack/echo/v4"
)

func (h *Handler) Register(e *echo.Echo) {
	withEmployee := []echo.MiddlewareFunc{
		middlewares.RequireAuth(),
		middlewares.LoadEmployee(h.app.Db),
		middlewares.LoadPermissions(h.app.Db, h.app.Services.Permissions),
	}

	withOwner := []echo.MiddlewareFunc{
		middlewares.RequireAuth(),
		middlewares.RequireOwner(h.app.Db),
	}

	e.DELETE("/admin/companies/delete/:id", h.deleteCompany, withOwner...)
	e.PUT("/admin/companies/update/:id", h.updateCompany, withOwner...)

	e.POST("/admin/companies/invites/sendInvite", h.sendCompanyInvite, withOwner...)
	e.POST("/admin/companies/invites/resend/:id", h.resendCompanyInvite, withOwner...)
	e.DELETE("/admin/companies/invites/delete/:id", h.deleteCompanyInvite, withOwner...)
	e.PATCH("/admin/companies/invites/reactivate/:id", h.reactivateCompanyInvite, withOwner...)

	e.POST("/auth/create", h.authSignUp)
	e.POST("/auth/processLogin", h.processLogin)

	e.POST("/app/suppliers/create/new", h.createSupplier, withEmployee...)
	e.PUT("/app/suppliers/information/edit/:name", h.updateSupplier, withEmployee...)
	e.POST("/app/suppliers/create/product/:id", h.createProduct, withEmployee...)
	e.DELETE("/app/suppliers/delete/product/:supplierID/:productID", h.deleteProduct, withEmployee...)

	e.POST("/app/order/send", h.sendOrder, withEmployee...)
}
