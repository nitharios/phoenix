package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	animal struct {
		Name        string  `json:"name"`
		Type        string  `json:"type"`
		Size        string  `json:"size"`
		Breed       string  `json:"breed"`
		Age         string  `json:"age"`
		Picture     string  `json:"pictureURL"`
		Description string  `json:"description"`
		Shelter     shelter `json:"shelter"`
	}
)

// GetAnimals provides a list of animals
func GetAnimals(c echo.Context) error {
	sID := c.Param("sid")
	body := `Animals returned from Shelter ` + sID

	return c.HTML(http.StatusOK, body)
}

// AddAnimal creates a new animal for the shelter list
func AddAnimal(c echo.Context) error {
	body := `Animal created`

	return c.HTML(http.StatusOK, body)
}

// GetAnimal provides details of an animal
func GetAnimal(c echo.Context) error {
	anID := c.Param("anid")
	body := `Animal ` + anID

	return c.HTML(http.StatusOK, body)
}

// UpdateAnimal updates a current animal on the shelter list
func UpdateAnimal(c echo.Context) error {
	anID := c.Param("anid")
	body := `Animal ` + anID + ` updated`

	return c.HTML(http.StatusOK, body)
}

// RemoveAnimal removes an animal from the shelter list
func RemoveAnimal(c echo.Context) error {
	anID := c.Param("anid")
	body := `Animal ` + anID + ` removed`

	return c.HTML(http.StatusOK, body)
}
