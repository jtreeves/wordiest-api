package controller

import (
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
	c.PrimitiveController.RegisterRoutes()
}
