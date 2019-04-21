package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	animal struct {
		ID          string  `json:"animal_id"`
		Name        string  `json:"name"`
		URL         string  `json:"url"`
		PictureURL  string  `json:"picture_url"`
		Type        string  `json:"type"`
		Size        string  `json:"size"`
		Breed       string  `json:"breed"`
		Age         string  `json:"age"`
		Description string  `json:"description"`
		Adopted     bool    `json:"adopted"`
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
	a := &animal{}
	if err := c.Bind(a); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, a)
}

// GetAnimal provides details of an animal
func GetAnimal(c echo.Context) error {
	anID := c.Param("anid")
	body := `Animal ` + anID

	return c.HTML(http.StatusOK, body)
}

// UpdateAnimal updates a current animal on the shelter list
func UpdateAnimal(c echo.Context) error {
	a := &animal{}
	if err := c.Bind(a); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, a)
}

// RemoveAnimal removes an animal from the shelter list
func RemoveAnimal(c echo.Context) error {
	anID := c.Param("anid")
	body := `Animal ` + anID + ` removed`

	return c.HTML(http.StatusOK, body)
}
