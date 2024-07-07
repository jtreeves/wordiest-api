package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndexController struct {
	*PrimitiveController
}

func NewIndexController(e *echo.Echo) *IndexController {
	g := e.Group("/")
	b := NewPrimitiveController(g)

	return &IndexController{PrimitiveController: b}
}

func (c *IndexController) RegisterRoutes() {
	c.Group.GET("", c.Get)
}

func (c *IndexController) Get(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello, World!")
}
