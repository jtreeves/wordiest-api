package main

import (
	_ "github.com/jtreeves/wordiest-api/docs"
	"github.com/jtreeves/wordiest-api/pkg/controller"
	"github.com/jtreeves/wordiest-api/pkg/service"
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

	ws := service.NewWordsService()

	index := controller.NewIndexController(e)
	docs := controller.NewDocsController(e)
	api := controller.NewApiController(e)
	words := controller.NewWordsController(api.Group, ws)

	index.RegisterRoutes()
	docs.RegisterRoutes()
	api.RegisterRoutes()
	words.RegisterRoutes()

	e.Logger.Fatal(e.Start(":1234"))
}
