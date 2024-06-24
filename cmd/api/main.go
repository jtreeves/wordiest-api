package main

import (
	"net/http"

	"github.com/jtreeves/wordiest-api/pkg/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/words", func(c echo.Context) error {
		return handler.GetWord(c)
	})

	e.Logger.Fatal(e.Start(":1234"))
}
