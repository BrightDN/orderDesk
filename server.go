package main

import (
	"database/sql"
	"html/template"
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

	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	template.Must(tmpl.ParseGlob("templates/components/*.html"))
	t := &configs.Template{
		Templates: tmpl,
	}

	e := echo.New()
	e.Renderer = t

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())
	e.Use(middlewares.ChangeMethod())
	e.Use(session.Middleware(sessions.NewCookieStore(
		[]byte(os.Getenv("SESSION_AUTH_KEY")),
		[]byte(os.Getenv("SESSION_ENCRYPT_KEY")),
	)))

	protected := e.Group("")
	protected.Use(middlewares.RequireAuth())
	protected.Use(middlewares.LoadUser(&cfg))

	e.Static("/assets", "assets")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "suppliers", nil)
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
