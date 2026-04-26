package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/brightDN/orderDesk/internal/database"
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
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Something went wrong loading the postgres db: %v", err)
	}

	dbQueries := database.New(db)

	cfg := configs.Config{
		Db:       dbQueries,
		Platform: os.Getenv("PLATFORM"),
	}

	cfg = cfg
	e := echo.New()
	e.Renderer = &configs.Template{}

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Use(middlewares.ChangeMethod())
	e.Use(session.Middleware(sessions.NewCookieStore(
		[]byte(os.Getenv("SESSION_AUTH_KEY")),
		[]byte(os.Getenv("SESSION_ENCRYPT_KEY")),
	)))

	// protected := e.Group("/dashboard")
	// protected.Use(middlewares.RequireAuth())
	// protected.Use(middlewares.LoadUser(&cfg))

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

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
