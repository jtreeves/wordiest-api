package controller

import (
	"github.com/jtreeves/wordiest-api/pkg/handler"
	"github.com/labstack/echo/v4"
)

func RegisterMainIndexRoute(api *echo.Group) {
	handler.RegisterIndexRoute(api)
}
