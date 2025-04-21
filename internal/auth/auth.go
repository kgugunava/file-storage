package auth

import (
	"github.com/kgugunava/file-storage/internal/models"
	"github.com/kgugunava/file-storage/internal/database"
	"github.com/labstack/echo/v4"
	"github.com/golang-jwt/jwt/v5"
	"time"
	"fmt"
	"net/http"
)

type AuthRequest struct {
	Login string `json:"login"`
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


func Authenticate(c echo.Context, db database.Database) error {
	request := new(AuthRequest)
	c.Bind(request)
	currentUser := models.User{Login: request.Login, Password: request.Password}
	if db.CheckAuthRequest(currentUser) == 1 {
		n := currentUser
		fmt.Print(n.Login)
		fmt.Print(n.Password)
		return c.NoContent(http.StatusOK)
	}
	return c.NoContent(http.StatusForbidden)
}