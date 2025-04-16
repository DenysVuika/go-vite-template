package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterRoutes(g *echo.Group) {
	g.GET("/books", getBooks)
	g.POST("/books", createBook)
}

func getBooks(c echo.Context) error {
	// Sample response
	return c.JSON(http.StatusOK, []string{"Book A", "Book B"})
}

func createBook(c echo.Context) error {
	// Placeholder
	return c.String(http.StatusCreated, "Book created!")
}
