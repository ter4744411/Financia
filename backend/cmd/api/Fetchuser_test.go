package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestFetchClientHandler(t *testing.T) {
	e := echo.New()

	// Create a new HTTP request with the idnumber parameter
	req := httptest.NewRequest(http.MethodGet, "/userdashboard", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/userdashboard/:idnumber")
	c.SetParamNames("idnumber")
	c.SetParamValues("1")

	// Call the handler function
	if assert.NoError(t, fetchUserHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expected := `{"idnumber":"1","username":"username-user1","role":"client"}`
		assert.JSONEq(t, expected, rec.Body.String())
	}
}

func TestFetchAdminHandler(t *testing.T) {
	e := echo.New()

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/admindashboard", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Mock user context with email
	user := &jwtCustomClaims{
		Email:    "admin1@gmail.com",
		Role:     "Admin",
		Idnumber: "1",
	}
	c.Set("user", user)

	// Call the handler function
	if assert.NoError(t, fetchAdminHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expected := `{"username":"username-admin1","role":"Admin"}`
		assert.JSONEq(t, expected, rec.Body.String())
	}
}
