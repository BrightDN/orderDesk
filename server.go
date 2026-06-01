package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/brightDN/orderDesk/internal/app"
	"github.com/brightDN/orderDesk/internal/configs"
	"github.com/brightDN/orderDesk/internal/database"
	"github.com/brightDN/orderDesk/internal/http/handlers"
	"github.com/brightDN/orderDesk/internal/http/routing"
	"github.com/brightDN/orderDesk/internal/middlewares"
	"github.com/brightDN/orderDesk/internal/services"
	"github.com/brightDN/orderDesk/internal/services/mailer"
	"github.com/labstack/echo/v4"

	_ "github.com/lib/pq"
)

func main() {
	cfg := configs.LoadConfigs()

	db, err := sql.Open(cfg.Db.Driver, cfg.Db.Url)
	if err != nil {
		log.Fatalf("failed to load the %s database: %v", cfg.Db.Driver, err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	ms, err := mailer.NewMailerService(cfg.Mail)
	if err != nil {
		log.Fatalf("Creating mailerclient failed: %v", err)
	}
	defer ms.Close()

	e := echo.New()
	defer e.Close()
	e.Renderer = &configs.Template{
		Identity: cfg.Identity,
	}
	e.HTTPErrorHandler = configs.HTTPErrorHandler
	e.Static("/assets", "assets")

	middlewares.Register(e, cfg)
	services := services.NewServices(dbQueries, db, ms, &cfg.Identity)
	app := app.New(services, dbQueries, cfg)

	n := routing.NewNav(dbQueries, &app)
	n.Register(e)

	h := handlers.NewHandler(&app)
	h.Register(e)

	httpPort := os.Getenv("PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}
