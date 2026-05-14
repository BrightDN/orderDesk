package main

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/brightDN/orderDesk/internal/app"
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/handlers"
	"github.com/brightDN/orderDesk/internal/mailer"
	"github.com/brightDN/orderDesk/internal/middlewares"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	dbURL := os.Getenv("DB_URL")
	sessionAuthKey := []byte(os.Getenv("SESSION_AUTH_KEY"))
	sessionEncryptKey, err := sessionEncryptionKey(os.Getenv("SESSION_ENCRYPT_KEY"))
	if err != nil {
		log.Fatalf("Invalid SESSION_ENCRYPT_KEY: %v", err)
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Something went wrong loading the postgres db: %v", err)
	}

	dbQueries := database.New(db)
	mc, err := mailer.NewClient("smtp-relay.brevo.com", 587, os.Getenv("MAILER_USER"), os.Getenv("MAILER_SECRET"))

	cfg := configs.Config{
		Platform:    os.Getenv("PLATFORM"),
		MailAccount: os.Getenv("MAILER_MAIL"),
	}

	e := echo.New()
	e.Renderer = &configs.Template{}
	e.HTTPErrorHandler = configs.HTTPErrorHandler
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:_csrf,header:" + echo.HeaderXCSRFToken,
	}))
	e.Use(middlewares.ChangeMethod())
	e.Use(session.Middleware(sessions.NewCookieStore(
		sessionAuthKey,
		sessionEncryptKey,
	)))

	app := app.New(dbQueries, cfg, mc, "OrderDesk")
	h := handlers.NewHandler(&app)

	// protected := e.Group("/dashboard")
	// protected.Use(middlewares.RequireAuth())
	// protected.Use(middlewares.LoadUser(app.Db))

	e.Static("/assets", "assets")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "newOrder", map[string]any{
			"Page": "new order",
		})
	})

	e.GET("/dashboard/suppliers", func(c echo.Context) error {
		return c.Render(http.StatusOK, "suppliers", map[string]any{
			"Page": "suppliers",
		})
	})

	e.GET("/dashboard/neworder", func(c echo.Context) error {
		return c.Render(http.StatusOK, "newOrder", map[string]any{
			"Page": "new order",
		})
	})

	e.GET("/dashboard/history", func(c echo.Context) error {
		return c.Render(http.StatusOK, "orderHistory", map[string]any{
			"Page": "history",
		})
	})

	e.GET("/dashboard/settings", func(c echo.Context) error {
		return c.Render(http.StatusOK, "companySettings", map[string]any{
			"Page": "company settings",
		})
	})

	e.GET("/support/contact", func(c echo.Context) error {
		return c.Render(http.StatusOK, "contactPage", nil)
	})

	e.GET("/auth/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login", nil)
	})

	e.GET("/auth/signup", func(c echo.Context) error {
		return c.Render(http.StatusOK, "signup", nil)
	})

	e.GET("/settings/companies", func(c echo.Context) error {
		return c.Render(http.StatusOK, "companies", nil)
	})

	e.GET("/settings/user", func(c echo.Context) error {
		return c.Render(http.StatusOK, "userSettings", nil)
	})

	e.GET("/admin/companies/invites", h.NavAdminNewCompany)
	e.GET("/admin/companies/list", h.NavAdminCompanyList)
	e.POST("/admin/companies/invites/sendInvite", h.SendCompanyInvite)
	e.POST("/admin/companies/invites/resend/:id", h.ResendCompanyInvite)
	e.DELETE("/admin/companies/invites/delete/:id", h.DeleteCompanyInvite)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

func sessionEncryptionKey(value string) ([]byte, error) {
	if len(value) == 32 || len(value) == 24 || len(value) == 16 {
		return []byte(value), nil
	}

	key, err := hex.DecodeString(value)
	if err != nil {
		return nil, err
	}
	if len(key) != 32 && len(key) != 24 && len(key) != 16 {
		return nil, fmt.Errorf("must be 16, 24, or 32 bytes after hex decoding; got %d bytes", len(key))
	}

	return key, nil
}
