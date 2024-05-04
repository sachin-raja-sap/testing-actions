package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

var users []User

func main() {
	e := echo.New()
	handleRoutes(e)
	e.Logger.Fatal(e.Start(":8000"))
}

func handleRoutes(e *echo.Echo) {
	e.GET("/health", Health)
	e.GET("/version", GetVersion(echo.Version))
	e.POST("/create", CreateUser)
	e.GET("/users", GetUsers)
}

func Health(c echo.Context) error {
	return c.JSON(http.StatusOK, "system is healthy")
}

func GetVersion(version string) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(
			http.StatusOK,
			&struct {
				Application string `json:"application"`
				Version     string `json:"version"`
			}{
				Application: "testing-actions",
				Version:     version,
			})
	}
}

func CreateUser(c echo.Context) error {
	reqBody := new(User)
	if err := c.Bind(reqBody); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	users = append(users, *reqBody)
	return c.JSON(http.StatusCreated, users)
}

func GetUsers(c echo.Context) error {
	if len(users) == 0 {
		return c.JSON(http.StatusNotFound, "No users found, please create a user first")
	}
	return c.JSON(http.StatusOK, users)
}
