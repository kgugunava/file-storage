package server

import (
	"file-storage/internal/database"
	"file-storage/internal/register"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RunServer(db database.Database) {
	// cfg := config.Config{}
	// config.Load(&cfg)
	// fmt.Print(cfg.Port)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Main Page")
	})
	registerHandler := func(c echo.Context) error {
		return register.RegisterUser(c, db)
	}
	e.POST("/register", registerHandler)
	e.Logger.Fatal(e.Start(":8010"))
}