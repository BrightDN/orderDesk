package middlewares

import (
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(e *echo.Echo, cfg configs.Config) {
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:_csrf,header:" + echo.HeaderXCSRFToken,
	}))
	e.Use(changeMethod())
	e.Use(session.Middleware(sessions.NewCookieStore(
		cfg.Session.SessionAuthKey,
		cfg.Session.SessionEncryptionKey,
	)))

	// protected := e.Group("/dashboard")
	// protected.Use(middlewares.RequireAuth())
	// protected.Use(middlewares.loadUser(app.Db))
}
