package controller

import (
	"github.com/jtreeves/wordiest-api/pkg/service"
	"github.com/labstack/echo/v4"
)

type WordsController struct {
	*CompositeController
}

func NewWordsController(e *echo.Group, s service.Service) *WordsController {
	g := e.Group("/words")
	b := NewCompositeController(g, s)

	return &WordsController{CompositeController: b}
}

func (c *WordsController) RegisterRoutes() {
	c.CompositeController.RegisterRoutes()
}

// @Summary Get single word
// @Description Returns a random word from the database
// @Tags words
// @Accept json
// @Produce plain
// @Success 200	{string} string "word"
// @Router /words [get]
func (c *WordsController) Get(ctx echo.Context) error {
	return c.CompositeController.Get(ctx)
}
