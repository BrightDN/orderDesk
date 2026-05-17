package handlers

import (
	"errors"

	"github.com/brightDN/orderDesk/internal/app"
)

type Handler struct {
	App *app.App
}

var ErrUnexpectedValue = errors.New("Unexpected value, action failed")
var ErrInternalError = errors.New("Something went wrong and we could not complete your request")

func NewHandler(app *app.App) *Handler {
	return &Handler{
		App: app,
	}
}
