package server

import (
	"log"
	"net/http"
	"github.com/labstack/echo/v4"
	"file-storage/internal/database"
	"file-storage/internal/login"
)

func RunServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Main Page")
	})
	e.POST("/login", loginUser)
	e.Logger.Fatal(e.Start(":8010"))
}

func loginUser(c echo.Context) error {
	user := c.FormValue("username")
	password := c.FormValue("password")
	currentUser := login.User{Login: user, Password: password}
	// fmt.Println(user, password)
	conn, err := database.ConnectToDatabase()
	if err != nil {
        log.Fatalf("Error : %v", err)
    }
	database.AddUserToDatabase(*conn, currentUser)
	return c.String(http.StatusOK, user)
}