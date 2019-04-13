package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// New creates a new server instance
func New(HOST string) *echo.Echo {
	e := echo.New()

	fmt.Println("Hosting on ", HOST)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", Home)

	return e
}
