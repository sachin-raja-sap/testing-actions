package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, Health(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `"system is healthy"`, rec.Body.String())
}

func TestGetVersion(t *testing.T) {
	e := echo.New()
	handler := GetVersion("1.0.0")

	req := httptest.NewRequest(http.MethodGet, "/version", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, handler(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	expectedJSON := `{"application":"testing-actions","version":"1.0.0"}`
	assert.JSONEq(t, expectedJSON, rec.Body.String())
}

func TestCreateUser(t *testing.T) {
	e := echo.New()

	userJSON := `{"id":"1", "username":"testuser", "email":"test@example.com"}`
	req := httptest.NewRequest(http.MethodPost, "/create", bytes.NewBufferString(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, CreateUser(c))
	assert.Equal(t, http.StatusCreated, rec.Code)
	// Check if the body contains the user data (assuming JSON output includes user array)
	assert.Contains(t, rec.Body.String(), "testuser")
	// Reset users for consistent test results
	users = []User{}
}

func TestGetUsers(t *testing.T) {
	e := echo.New()

	// Test with no users initially
	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	assert.NoError(t, GetUsers(c))
	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Contains(t, rec.Body.String(), "No users found, please create a user first")

	// Add a user and test again
	users = append(users, User{ID: "1", Username: "testuser", Email: "test@example.com"},
		User{ID: "2", Username: "testuser2", Email: "test@example.com"})

	req = httptest.NewRequest(http.MethodGet, "/users", nil)
	rec = httptest.NewRecorder()
	c = e.NewContext(req, rec)

	assert.NoError(t, GetUsers(c))
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "testuser")
	assert.Contains(t, rec.Body.String(), "testuser2")

	// Reset users for consistent test results
	users = []User{}
}
