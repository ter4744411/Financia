package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGetUserInfoHandler(t *testing.T) {
	e := echo.New()

	// Create a new HTTP request
	req := httptest.NewRequest(http.MethodGet, "/userinfo", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Mock user context with email
	user := &jwtCustomClaims{
		Email:    "user2@gmail.com",
		Role:     "client",
		Idnumber: "2",
	}
	c.Set("user", user)

	// Call the handler function
	if assert.NoError(t, getUserInfoHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		expected := `{"idnumber":"2","username":"username-user2","email":"user2@gmail.com","role":"client"}`
		assert.JSONEq(t, expected, rec.Body.String())
	}
}
