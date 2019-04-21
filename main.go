package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	m "github.com/nitharios/phoenix/middleware"
)

func main() {
	// Set environment variables
	ENV := os.Getenv("ENV")
	HOST := os.Getenv("HOST")
	PORT := ":" + os.Getenv("PORT")
	TOKEN := os.Getenv("TOKEN")

	// Create new client
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	// Login
	v1 := e.Group("/v1")
	v1.POST("/login", func(c echo.Context) error {
		return m.Login(c, TOKEN)
	})
	v1.Use(middleware.JWT([]byte(TOKEN)))
	v1 = GenerateRoutes(v1)

	// Set environment conditions
	if ENV != "PROD" {
		fmt.Println("Starting test server on", HOST, ":8000")
		e.Logger.Fatal(e.Start(":8000"))

	} else {
		fmt.Println("Starting production server on", PORT)
		e.Logger.Fatal(e.Start(PORT))
	}
}
