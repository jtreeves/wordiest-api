package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type APIController struct {
	*PrimitiveController
}

func NewAPIController(e *echo.Echo) *APIController {
	g := e.Group("/api")
	b := NewPrimitiveController(g)

	return &APIController{PrimitiveController: b}
}

func (c *APIController) RegisterRoutes() {
	c.Group.GET("", c.Get)
}

func (c *APIController) Get(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "API")
}
