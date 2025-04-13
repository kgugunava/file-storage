package register

import (
	"file-storage/internal/database"
	"log"
	"github.com/labstack/echo/v4"
	"net/http"
	"file-storage/internal/userstruct"
)

type RegisterRequest struct {
	Login string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(c echo.Context) error{
	request := new(RegisterRequest)
	if err := c.Bind(request); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	currentUser := userstruct.User{Login: request.Login, Password: request.Password}
	conn, err := database.ConnectToDatabase()
	if err != nil {
        log.Fatalf("Error : %v", err)
    }
	database.AddUserToDatabase(*conn, currentUser)
	return c.NoContent(http.StatusOK)
}