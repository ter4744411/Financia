package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

func getUserInfoHandler(c echo.Context) error {
	user, ok := c.Get("user").(*jwtCustomClaims)
	if !ok {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token from get userinfohandler")
	}
	email := user.Email

	fmt.Printf("Fetching info for user with email: %s\n", email)

	userpath := `C:\Users\User\OneDrive\Desktop\financia\frontend\users.csv`
	userrecords, err := readCSVFile(userpath)
	if err != nil {
		fmt.Printf("Error reading CSV file: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Read csv error")
	}

	for _, record := range userrecords {
		if record[2] == email {
			userInfo := map[string]string{
				"idnumber": record[0],
				"username": record[1],
				"email":    record[2],
				"role":     record[4],
			}
			fmt.Printf("User found: %v\n", userInfo)
			return c.JSON(http.StatusOK, userInfo)
		}
	}

	adminpath := `C:\Users\User\OneDrive\Desktop\financia\frontend\admins.csv`
	adminrecords, err := readCSVFile(adminpath)
	if err != nil {
		fmt.Printf("Error reading CSV file: %v\n", err)
		return echo.NewHTTPError(http.StatusInternalServerError, "Read csv error")
	}

	for _, record := range adminrecords {
		if record[2] == email {
			userInfo := map[string]string{
				"idnumber": record[0],
				"username": record[1],
				"email":    record[2],
				"role":     record[4],
			}
			fmt.Printf("User found: %v\n", userInfo)
			return c.JSON(http.StatusOK, userInfo)
		}
	}

	fmt.Printf("User not found: %s\n", email)
	return echo.NewHTTPError(http.StatusNotFound, "User not found")
}
