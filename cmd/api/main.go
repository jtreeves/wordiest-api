package main

import (
	"net/http"

	_ "github.com/jtreeves/wordiest-api/docs"
	"github.com/jtreeves/wordiest-api/pkg/handlers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Wordiest API
// @version 1.0
// @description A simple API for receiving word data
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:1234
// @BasePath /api
func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/docs/*", echoSwagger.WrapHandler)

	api := e.Group("/api")

	handlers.RegisterWordsRoute(api)

	e.Logger.Fatal(e.Start(":1234"))
}
