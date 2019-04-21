package main

import (
	"github.com/labstack/echo/v4"

	m "github.com/nitharios/phoenix/middleware"
)

// GenerateRoutes builds the available routes for the api
func GenerateRoutes(g *echo.Group) *echo.Group {
	// Shelters routes
	g.GET("/shelters", m.GetShelters)
	g.POST("/shelters", m.AddShelter)
	g.GET("/shelters/:sid", m.GetShelter)
	g.PATCH("/shelters/:sid", m.UpdateShelter)
	g.DELETE("/shelters/:sid", m.RemoveShelter)
	sURI := "/shelters/:sid"

	// Animals routes
	g.GET(sURI+"/animals", m.GetAnimals)
	g.POST(sURI+"/animals", m.AddAnimal)
	g.GET(sURI+"/animals/:anid", m.GetAnimal)
	g.PATCH(sURI+"/animals/:anid", m.UpdateAnimal)
	g.DELETE(sURI+"/animals/:anid", m.RemoveAnimal)

	return g
}
