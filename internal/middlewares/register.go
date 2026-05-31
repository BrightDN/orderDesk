package middlewares

import (
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(e *echo.Echo, cfg configs.Config, q *database.Queries) {
	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: func(c echo.Context) string {
			m := c.FormValue("_method")
			return m
		},
	}))

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:_csrf,header:" + echo.HeaderXCSRFToken,
	}))

	e.Use(session.Middleware(sessions.NewCookieStore(
		cfg.Session.SessionAuthKey,
		cfg.Session.SessionEncryptionKey,
	)))

	protected := e.Group("/app")
	protected.Use(requireAuth())
	protected.Use(loadEmployee(q))
}
