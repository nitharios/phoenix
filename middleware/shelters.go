package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	shelter struct {
		ID            string   `json:"shelter_id"`
		Name          string   `json:"name"`
		AdoptionsPage string   `json:"adoptions_page"`
		Homepage      string   `json:"homepage"`
		Email         string   `json:"email"`
		PhoneNumber   string   `json:"phone_number"`
		Location      location `json:"location"`
	}

	// Location format
	location struct {
		Street  string `json:"street"`
		City    string `json:"city"`
		State   string `json:"state"`
		ZipCode string `json:"zip_code"`
	}
)

// GetShelters provides a list of animal shelters
func GetShelters(c echo.Context) error {
	body := `Shelters returned`

	return c.HTML(http.StatusOK, body)
}

// AddShelter creates a new animal shelter for the list
func AddShelter(c echo.Context) error {
	s := &shelter{}
	if err := c.Bind(s); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, s)
}

// GetShelter provides details of an animal shelter
func GetShelter(c echo.Context) error {
	sID := c.Param("sid")
	body := `Shelter ` + sID

	return c.HTML(http.StatusOK, body)
}

// UpdateShelter updates a current animal shelter on the list
func UpdateShelter(c echo.Context) error {
	s := &shelter{}
	if err := c.Bind(s); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, s)
}

// RemoveShelter removes an animal shelter from the list
func RemoveShelter(c echo.Context) error {
	sID := c.Param("sid")
	body := `Shelter ` + sID + ` removed`

	return c.HTML(http.StatusOK, body)
}
