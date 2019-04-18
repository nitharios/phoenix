package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	shelter struct {
		Name        string   `json:"shelterName"`
		Website     string   `json:"website"`
		Email       string   `json:"email"`
		PhoneNumber string   `json:"phoneNumber"`
		Location    location `json:"location"`
	}

	// Location format
	location struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		State   string `json:"state"`
		ZipCode string `json:"zipCode"`
	}
)

// GetShelters provides a list of animal shelters
func GetShelters(c echo.Context) error {
	body := `Shelters returned`

	return c.HTML(http.StatusOK, body)
}

// AddShelter creates a new animal shelter for the list
func AddShelter(c echo.Context) error {
	body := `Shelter created`

	return c.HTML(http.StatusOK, body)
}

// GetShelter provides details of an animal shelter
func GetShelter(c echo.Context) error {
	sID := c.Param("sid")
	body := `Shelter ` + sID

	return c.HTML(http.StatusOK, body)
}

// UpdateShelter updates a current animal shelter on the list
func UpdateShelter(c echo.Context) error {
	sID := c.Param("sid")
	body := `Shelter ` + sID + ` updated`

	return c.HTML(http.StatusOK, body)
}

// RemoveShelter removes an animal shelter from the list
func RemoveShelter(c echo.Context) error {
	sID := c.Param("sid")
	body := `Shelter ` + sID + ` removed`

	return c.HTML(http.StatusOK, body)
}
