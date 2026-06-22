package middlewares

import (
	"net/http"

	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Register(e *echo.Echo, cfg configs.Config) {
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

	store := sessions.NewCookieStore(
		cfg.Session.SessionAuthKey,
		cfg.Session.SessionEncryptionKey,
	)
	store.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	e.Use(session.Middleware(store))
}
