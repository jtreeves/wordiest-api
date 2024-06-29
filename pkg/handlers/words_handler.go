// Handle API calls to various endpoints
package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterWordsRoute(api *echo.Group) {
	words := api.Group("/words")
	words.GET("", getWord)
}

// @Summary Get single word
// @Description Returns a random word from the database
// @Tags words
// @Accept json
// @Produce plain
// @Success 200	{string} string "word"
// @Router /words [get]
func getWord(c echo.Context) error {
	return c.String(http.StatusOK, "word")
}
