package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func runServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Main Page")
	})
	t := &Template{
		templates: template.Must(template.ParseGlob("templates/*.html")),
	}
	e.Renderer = t
	e.GET("/login", login)
	e.Logger.Fatal(e.Start(":8010"))
}

func login(c echo.Context) error {
	userLogin := c.FormValue("username")
	userPassword := c.FormValue("password")
	fmt.Println(userLogin)
	fmt.Println(userPassword)
 	return c.Render(http.StatusOK, "loginPage", "")
}