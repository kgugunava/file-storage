package server

import (
	"github.com/kgugunava/file-storage/internal/database"
	"file-storage/internal/register"
	"net/http"
	"file-storage/internal/auth"
	"file-storage/internal/config"
	"github.com/labstack/echo/v4"
	"fmt"
)

func RunServer(db database.Database) {
	cfg, _ := config.Load()
	fmt.Print("AAAAAA")
	fmt.Print(cfg)
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Main Page")
	})
	registerHandler := func(c echo.Context) error {
		return register.RegisterUser(c, db)
	}
	authHandler := func(c echo.Context) error {
		return auth.Authenticate(c, db)
	}
	e.POST("/register", registerHandler)
	e.POST("/auth", authHandler)
	e.Logger.Fatal(e.Start(cfg.Port))
}