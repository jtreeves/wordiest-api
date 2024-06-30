package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterIndexRoute(index *echo.Group) {
	index.GET("", getIndex)
}

func getIndex(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
