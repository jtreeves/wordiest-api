package main

import (
	"net/http"

	_ "github.com/jtreeves/wordiest-api/docs"
	"github.com/jtreeves/wordiest-api/pkg/handlers"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

//	@title			Echo Swagger Example API
//	@version		1.0
//	@description	This is a sample server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:1234
//	@BasePath	/api/v1

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/docs/*", echoSwagger.WrapHandler)

	e.GET("/words", func(c echo.Context) error {
		return handlers.GetWord(c)
	})

	e.Logger.Fatal(e.Start(":1234"))
}
