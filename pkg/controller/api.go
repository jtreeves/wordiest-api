package controller

import (
	"net/http"

	"github.com/jtreeves/wordiest-api/pkg/service"
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

	ws := service.NewWordsService()
	words := NewWordsController(c.Group, ws)

	words.RegisterRoutes()
}

func (c *APIController) Get(ctx echo.Context) error {
	return ctx.JSON(http.StatusBadRequest, "Invalid request")
}
