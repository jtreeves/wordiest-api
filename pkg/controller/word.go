package controller

import (
	"net/http"

	"github.com/jtreeves/wordiest-api/pkg/service"
	"github.com/labstack/echo/v4"
)

type WordController struct {
	*CompositeController
}

func NewWordController(e *echo.Group, s service.Service) *WordController {
	g := e.Group("/word")
	b := NewCompositeController(g, s)

	return &WordController{CompositeController: b}
}

func (c *WordController) RegisterRoutes() {
	c.Group.GET("", c.Get)
}

// @Summary Get single word
// @Description Returns a random word from the database
// @Tags word
// @Accept json
// @Produce plain
// @Success 200	{string} string "word"
// @Router /word [get]
func (c *WordController) Get(ctx echo.Context) error {
	result := c.Service.Do()

	return ctx.JSON(http.StatusOK, result)
}
