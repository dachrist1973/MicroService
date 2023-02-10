package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    string `json: "id"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

func main() {
	e := echo.New()
	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)
	e.POST("/testPost/:id", testPostHandler)
	e.Logger.Fatal(e.Start(":8080"))

}

// Handler
func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func testPostHandler(c echo.Context) error {
	u := new(User)
	id := c.Param("id")

	c.Bind(u)
	u.Id = id
	return c.JSON(http.StatusOK, u)

}
