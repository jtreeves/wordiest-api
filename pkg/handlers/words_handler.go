// Handle API calls to various endpoints
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// @Summary Get single word
// @Description Get single word
// @Tags word
// @Accept json
// @Produce plain
// @Success 200	{string} string "word"
// @Router /words [get]
func GetWord(c echo.Context) error {
	return c.String(http.StatusOK, "word")
}
