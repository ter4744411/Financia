package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUserHandler_Valid(t *testing.T) {

}

func TestRegisterUserHandler_Invalid(t *testing.T) {

	t.Run("email input is missing return StatusBadRequest", func(t *testing.T) {

		requestBody := []byte(`{"email": "", "password": "registerpassword", "username": "regis-username"}`)
		req := httptest.NewRequest(http.MethodPost, "/registerclient", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)
		if assert.NoError(t, loginUserHandler(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("password input is missing return StatusBadRequest", func(t *testing.T) {

		requestBody := []byte(`{"email": "test@gmail.com", "password": "", "username": "regis-username"}`)
		req := httptest.NewRequest(http.MethodPost, "/registerclient", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)
		if assert.NoError(t, loginUserHandler(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
	t.Run("username input is missing return StatusBadRequest", func(t *testing.T) {

		requestBody := []byte(`{"email": "test@gmail.com", "password": "registerpassword", "username": ""}`)
		req := httptest.NewRequest(http.MethodPost, "/registerclient", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)
		if assert.NoError(t, loginUserHandler(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("email input is invalid return StatusBadRequest", func(t *testing.T) {

		requestBody := []byte(`{"email": "test", "password": "registerpassword", "username": "regis-username"}`)
		req := httptest.NewRequest(http.MethodPost, "/registerclient", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)
		if assert.NoError(t, loginUserHandler(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("password less than 8 characters return StatusBadRequest", func(t *testing.T) {

		requestBody := []byte(`{"email": "test@gmail.com", "password": "regis", "username": "regis-username"}`)
		req := httptest.NewRequest(http.MethodPost, "/registerclient", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)
		if assert.NoError(t, loginUserHandler(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("username over 10 characters return StatusBadRequest", func(t *testing.T) {

		requestBody := []byte(`{"email": "test", "password": "registerusername", "username": "regis-username"}`)
		req := httptest.NewRequest(http.MethodPost, "/registerclient", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)
		if assert.NoError(t, loginUserHandler(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})

	t.Run("email is existed in database return StatusBadRequest", func(t *testing.T) {

		requestBody := []byte(`{"email": "user1@gmail.com", "password": "registerusername", "username": "regis-username"}`)
		req := httptest.NewRequest(http.MethodPost, "/registerclient", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		e := echo.New()
		c := e.NewContext(req, rec)
		if assert.NoError(t, loginUserHandler(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
