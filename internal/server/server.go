package server

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"file-storage/internal/register"
	// "file-storage/internal/auth"
)

func RunServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Main Page")
	})
	e.POST("/register", register.RegisterUser)
	// e.POST("/login", auth.Authenticate)
	e.Logger.Fatal(e.Start(":8010"))
}
