package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func fetchUserHandler(c echo.Context) error {

	type Userinfo struct {
		IDNumber string `json:"idnumber"`
		Username string `json:"username"`
		Role     string `json:"role"`
	}

	idnumber := c.Param("idnumber")
	if idnumber == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "idnumber is required")
	}

	// Path to the CSV file
	csvPath := `C:\Users\User\OneDrive\Desktop\financia\frontend\users.csv`

	// Open the CSV file
	file, err := os.Open(csvPath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error opening CSV file: %v", err))
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error reading CSV file: %v", err))
	}

	// Find the user by idnumber
	for i, record := range records {
		if i == 0 {
			// Skip the header row
			continue
		}
		if record[0] == idnumber {
			user := Userinfo{
				IDNumber: record[0],
				Username: record[1],
				Role:     record[4],
			}
			return c.JSON(http.StatusOK, user)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "User not found")
}

func fetchAdminHandler(c echo.Context) error {
	user, ok := c.Get("user").(*jwtCustomClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token from get userinfohandler")
	}
	email := user.Email
	type Admininfo struct {
		Username string `json:"username"`
		Role     string `json:"role"`
	}
	// Path to the CSV file
	csvPath := `C:\Users\User\OneDrive\Desktop\financia\frontend\admins.csv`

	// Open the CSV file
	file, err := os.Open(csvPath)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error opening CSV file: %v", err))
	}
	defer file.Close()

	// Read the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("Error reading CSV file: %v", err))
	}

	// Find the user by idnumber
	for i, record := range records {
		if i == 0 {
			// Skip the header row
			continue
		}
		if record[2] == email {
			user := Admininfo{
				Username: record[1],
				Role:     record[4],
			}
			return c.JSON(http.StatusOK, user)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "User not found")
}
