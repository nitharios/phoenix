package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Home is a basic return message for requests to the home page
func Home(c echo.Context) error {
	body := `<h1>Phoenix</h1>
		<h3>Hello World!</h3>`

	return c.HTML(http.StatusOK, body)
}
