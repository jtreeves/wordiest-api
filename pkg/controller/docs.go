package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type DocsController struct {
	*PrimitiveController
}

func NewDocsController(e *echo.Echo) *DocsController {
	g := e.Group("/docs")
	b := NewPrimitiveController(g)

	return &DocsController{PrimitiveController: b}
}

func (c *DocsController) RegisterRoutes() {
	c.PrimitiveController.RegisterRoutes()
	c.Group.GET("/*", echoSwagger.WrapHandler)
}

func (c *DocsController) Get(ctx echo.Context) error {
	return ctx.Redirect(http.StatusFound, "/docs/index.html")
}
