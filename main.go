package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Set environment variables
	ENV := os.Getenv("ENV")
	HOST := os.Getenv("HOST")
	PORT := ":" + os.Getenv("PORT")

	// Create echo client
	e := echo.New()

	// Set logging and recover functionality
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e = GenerateRoutes(e)

	// Set environment conditions
	if ENV != "PROD" {
		fmt.Println("Starting test server on", HOST, ":8000")
		e.Logger.Fatal(e.Start(":8000"))

	} else {
		fmt.Println("Starting production server on", PORT)
		e.Logger.Fatal(e.Start(PORT))
	}
}
