package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func runServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Main Page")
	})
	e.POST("/login", login)
	e.Logger.Fatal(e.Start(":8010"))
}

func login(c echo.Context) error {
	user := c.FormValue("username")
	password := c.FormValue("password")
	currentUser := User{Login: user, Password: password}
	// fmt.Println(user, password)
	conn, err := connectToDatabase()
	if err != nil {
        log.Fatalf("Error : %v", err)
    }
	addUserToDatabase(*conn, currentUser)
	return c.String(http.StatusOK, user)
}