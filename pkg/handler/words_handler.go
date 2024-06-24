// Handle API calls to various endpoints
package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Get a single word
func GetWord(c echo.Context) error {
	return c.String(http.StatusOK, "word")
}
