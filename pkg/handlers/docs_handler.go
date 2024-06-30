package handlers

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterDocsRoutes(docs *echo.Group) {
	docs.GET("", getLanding)
	docs.GET("/*", echoSwagger.WrapHandler)
}

func getLanding(c echo.Context) error {
	return c.Redirect(302, "/docs/index.html")
}
