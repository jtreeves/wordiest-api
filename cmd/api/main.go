package main

import (
	_ "github.com/jtreeves/wordiest-api/docs"
	"github.com/jtreeves/wordiest-api/pkg/handlers"
	"github.com/labstack/echo/v4"
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
	index := e.Group("/")
	docs := e.Group("/docs")
	api := e.Group("/api")

	handlers.RegisterIndexRoute(index)
	handlers.RegisterDocsRoutes(docs)
	handlers.RegisterApiRoutes(api)

	e.Logger.Fatal(e.Start(":1234"))
}
