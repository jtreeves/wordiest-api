package controller

import (
	"github.com/labstack/echo/v4"
)

type ApiController struct {
	*PrimitiveController
}

func NewApiController(e *echo.Echo) *ApiController {
	g := e.Group("/api")
	b := NewPrimitiveController(g)

	return &ApiController{PrimitiveController: b}
}

func (c *ApiController) RegisterRoutes() {
	c.PrimitiveController.RegisterRoutes()
}
