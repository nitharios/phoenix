package main

import (
	"github.com/labstack/echo/v4"

	m "github.com/nitharios/phoenix/middleware"
)

// GenerateRoutes builds the available routes for the api
func GenerateRoutes(e *echo.Echo) *echo.Echo {
	e.GET("/", m.Home)

	return e
}
