package register

import (
	"file-storage/internal/database"
	"github.com/labstack/echo/v4"
	"net/http"
	"file-storage/internal/models"
	"fmt"
)

type RegisterRequest struct {
	Login string `json:"login"`
	Password string `json:"password"`
}

func RegisterUser(c echo.Context, db database.Database) error{
	request := new(RegisterRequest)
	if err := c.Bind(request); err != nil {
		fmt.Print("bind error\n")
		return c.String(http.StatusBadRequest, "bad request")
	}
	currentUser := models.User{Login: request.Login, Password: request.Password}
	db.AddUserToDatabase(currentUser)
	return c.NoContent(http.StatusOK)
}