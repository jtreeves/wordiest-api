package controller

import (
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
