package main

import (
	"github.com/labstack/echo/v4"

	m "github.com/nitharios/phoenix/middleware"
)

// GenerateRoutes builds the available routes for the api
func GenerateRoutes(e *echo.Echo) *echo.Echo {
	// Shelters routes
	e.GET("/shelters", m.GetShelters)
	e.POST("/shelters", m.AddShelter)
	e.GET("/shelters/:sid", m.GetShelter)
	e.PATCH("/shelters/:sid", m.UpdateShelter)
	e.DELETE("/shelters/:sid", m.RemoveShelter)
	sURI := e.URI(m.GetShelter)

	// Animals routes
	e.GET(sURI+"/animals", m.GetAnimals)
	e.POST(sURI+"/animals", m.AddAnimal)
	e.GET(sURI+"/animals/:anid", m.GetAnimal)
	e.PATCH(sURI+"/animals/:anid", m.UpdateAnimal)
	e.DELETE(sURI+"/animals/:anid", m.RemoveAnimal)

	return e
}
