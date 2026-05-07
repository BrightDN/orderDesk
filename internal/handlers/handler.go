package handlers

import "github.com/brightDN/orderDesk/internal/app"

type Handler struct {
	App *app.App
}

func NewHandler(app *app.App) *Handler {
	return &Handler{
		App: app,
	}
}
