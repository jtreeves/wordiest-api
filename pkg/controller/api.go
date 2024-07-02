package controller

import (
	"github.com/jtreeves/wordiest-api/pkg/handler"
	"github.com/labstack/echo/v4"
)

func RegisterMainApiRoutes(api *echo.Group) {
	handler.RegisterWordsRoutes(api)
}
