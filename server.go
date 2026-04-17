package main

func main() {
        godotenv.Load()
        dbURL := os.Getenv("DB_URL")
        db, err := sql.Open("postgres", dbURL)
        if err != nil {
                log.Fatalf("Something went wrong loading the postgres db: %v", err)
        }

        dbQueries := database.New(db)

        cfg := configs.Config{
                Db:         dbQueries,
                SecretPass: os.Getenv("SECRET_PASS"),
        }
        t := &configs.Template{
                Templates: template.Must(template.ParseGlob("public/*.html")),
        }

        e := echo.New()
        e.Renderer = t

        e.Use(middleware.RequestLogger())
        e.Use(middleware.Recover())
        e.Use(middleware.CSRF())

        e.Static("/assets", "assets")

        e.GET("/", func(c echo.Context) error {
                projects, err := cfg.Db.GetAllProjects(c.Request().Context())
                if err != nil {
                        return c.JSON(http.StatusInternalServerError, struct{ Status string }{Status: "Something went wrong connecting to the database"})
                }

                groupedProjects := make(map[string][]database.Project)
                for _, p := range projects {
                        groupedProjects[p.ProjectType] = append(groupedProjects[p.ProjectType], p)
                }

                return c.Render(http.StatusOK, "index.html", map[string]any{
                        "Projects": groupedProjects,
                })
        })

        e.GET("/admin", routings.MiddlewareHandler(cfg, routings.HandleAdmin), middleware.BasicAuth(func(user, pass string, c echo.Context) (bool, error) {
                if user == "admin" && pass == cfg.SecretPass {
                        return true, nil
                }
                return false, nil
        }))

        e.POST("/admin/projecttypes/add", routings.MiddlewareHandler(cfg, routings.HandleAddProjectType))
        e.POST("/admin/projecttypes/delete", routings.MiddlewareHandler(cfg, routings.HandleDeleteProjectType))
        e.POST("/admin/projecttypes/update", routings.MiddlewareHandler(cfg, routings.HandleAlterProjectType))

        e.POST("/admin/projects/add", routings.MiddlewareHandler(cfg, routings.HandleAddProject))
        e.POST("/admin/projects/delete", routings.MiddlewareHandler(cfg, routings.HandleDeleteProject))
        e.POST("/admin/projects/edit", routings.MiddlewareHandler(cfg, routings.HandleEditProject))

        e.GET("/health", func(c echo.Context) error {
                return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
        })

        httpPort := os.Getenv("PORT")
        if httpPort == "" {
                httpPort = "8080"
        }

        e.Logger.Fatal(e.Start(":" + httpPort))
}
