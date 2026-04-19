package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/brightDN/orderDesk/internal/apiConfigs"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/joho/godotenv"
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

	cfg := apiConfigs.Config{
		Db: dbQueries,
	}

	tmpl := template.Must(template.ParseGlob("templates/*.html"))
	template.Must(tmpl.ParseGlob("templates/components/*.html"))
	t := &apiConfigs.Template{
		Templates: tmpl,
	}

	cfg = cfg

	e := echo.New()
	e.Renderer = t

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	e.Use(middleware.CSRF())

	e.Static("/assets", "assets")

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "home", nil)
	})

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
