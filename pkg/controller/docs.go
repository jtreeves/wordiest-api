package controller

import (
	"github.com/jtreeves/wordiest-api/pkg/handler"
	"github.com/labstack/echo/v4"
)

func RegisterMainDocsRoutes(docs *echo.Group) {
	handler.RegisterDocsRoutes(docs)
}
