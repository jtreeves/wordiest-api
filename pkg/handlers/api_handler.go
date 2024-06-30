// Handle API calls to various endpoints
package handlers

import (
	"github.com/labstack/echo/v4"
)

func RegisterApiRoutes(api *echo.Group) {
	RegisterWordsRoutes(api)
}
