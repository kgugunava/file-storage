package auth

import (
	"file-storage/internal/models"
	"file-storage/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"fmt"
)

type AuthRequest struct {
	Login string `json:"email"`
	Password string `json:"password"`
}

func generateJWT() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2025, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	tokenString, _ := token.SignedString("hmacSampleSecret")
	fmt.Print(tokenString) 
	return tokenString
}


func Authenticate(c echo.Context) error{
	request := new(AuthRequest)
	currentUser := userstruct.User{Login: request.Login, Password: request.Password}
	conn, err := database.ConnectToDatabase()
	if err != nil {
        log.Fatalf("Error : %v", err)
    }
	if database.UserInDatabase(conn, currentUser) == 1, nil {
		n := currentUser
		fmt.Print(n.Login)
		fmt.Print(n.Password)
		database.AddUserToDatabase(conn, currentUser)
	}
}