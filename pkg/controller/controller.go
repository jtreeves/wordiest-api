package controller

import (
	"net/http"

	"github.com/jtreeves/wordiest-api/pkg/service"
	"github.com/labstack/echo/v4"
)

type Controller interface {
	RegisterRoutes()
	Get(ctx echo.Context) error
}

type PrimitiveController struct {
	Group *echo.Group
}

type CompositeController struct {
	*PrimitiveController
	Service service.Service
}

func NewPrimitiveController(g *echo.Group) *PrimitiveController {
	return &PrimitiveController{Group: g}
}

func NewCompositeController(g *echo.Group, s service.Service) *CompositeController {
	p := NewPrimitiveController(g)

	return &CompositeController{PrimitiveController: p, Service: s}
}

func (c *PrimitiveController) RegisterRoutes() {
	c.Group.GET("", c.Get)
}

func (c *PrimitiveController) Get(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "")
}

func (c *CompositeController) Get(ctx echo.Context) error {
	result := c.Service.Do()

	return ctx.JSON(http.StatusOK, result)
}
