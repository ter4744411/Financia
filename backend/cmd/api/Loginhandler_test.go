// main_test.go
package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestLoginUserHandler_Valid(t *testing.T) {
	requestBody := []byte(`{"email": "user1@gmail.com", "password": "user1"}`)
	req := httptest.NewRequest(http.MethodPost, "/loginuser", bytes.NewBuffer(requestBody))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	err := loginUserHandler(c)
	t.Run("email and password existed in database should return statusOk", func(t *testing.T) {

		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
	})
	t.Run("email and password existed in database should return token", func(t *testing.T) {
		var responseMap map[string]interface{}
		err = json.Unmarshal(rec.Body.Bytes(), &responseMap)
		assert.NoError(t, err)

		// Check that the "token" key is present and not empty
		_, exists := responseMap["token"]
		assert.True(t, exists)
	})
}

func TestLoginUserHandler_Invalid(t *testing.T) {
	t.Run("email input is missing return StatusBadRequest", func(t *testing.T) {
		requestBody := []byte(`{"email": "", "password": "wrongpassword"}`)
		req := httptest.NewRequest(http.MethodPost, "/loginuser", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)
		if assert.NoError(t, loginUserHandler(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("password input is missing return StatusBadRequest", func(t *testing.T) {
		requestBody := []byte(`{"email": "invalid@example.com", "password": ""}`)
		req := httptest.NewRequest(http.MethodPost, "/loginuser", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)
		if assert.NoError(t, loginUserHandler(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("test username is not existed in database return StatusBadRequest", func(t *testing.T) {
		requestBody := []byte(`{"email": "invalid@example.com", "password": "wrongpassword"}`)
		req := httptest.NewRequest(http.MethodPost, "/loginuser", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)
		if assert.NoError(t, loginUserHandler(c)) {

			// Check the response status code
			assert.Equal(t, http.StatusBadRequest, rec.Code)

			expectedResponseBody := `{"message":"Incorrect Email or Password!"}`
			assert.JSONEq(t, expectedResponseBody, rec.Body.String())
		}
	})
}
