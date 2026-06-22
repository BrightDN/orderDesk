package handlers

import "github.com/labstack/echo/v4"

func (h *Handler) Register(e *echo.Echo) {
	e.DELETE("/admin/companies/delete/:id", h.deleteCompany)
	e.PUT("/admin/companies/update/:id", h.updateCompany)

	e.POST("/admin/companies/invites/sendInvite", h.sendCompanyInvite)
	e.POST("/admin/companies/invites/resend/:id", h.resendCompanyInvite)
	e.DELETE("/admin/companies/invites/delete/:id", h.deleteCompanyInvite)
	e.PATCH("/admin/companies/invites/reactivate/:id", h.reactivateCompanyInvite)

	e.POST("/auth/create", h.authSignUp)
	e.POST("/auth/processLogin", h.processLogin)

	e.POST("/app/suppliers/create/new", h.createSupplier)
	e.PUT("/app/suppliers/information/edit/:name", h.updateSupplier)
	e.POST("/app/suppliers/create/product/:id", h.createProduct)
	e.DELETE("/app/suppliers/delete/product/:supplierID/:productID", h.deleteProduct)

	e.POST("/app/order/send", h.sendOrder)
}
