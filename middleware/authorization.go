package middleware

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

type (
	credentials struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

// Login function
func Login(c echo.Context, at string) error {
	cred := &credentials{}
	if err := c.Bind(cred); err != nil {
		return err
	}

	if cred.Username == "testuser" || cred.Password == "password" {
		return echo.ErrUnauthorized
	}

	// Create token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":  cred.Username,
		"admin": true,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	})

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(at))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}
