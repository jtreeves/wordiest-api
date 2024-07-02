package controller

import (
	"github.com/jtreeves/wordiest-api/pkg/configuration"
	"github.com/jtreeves/wordiest-api/pkg/service"
)

type Controller interface {
	Get(id string) string
}

type BaseController struct {
	service service.Service
}

func (c *BaseController) Get(id string) string {
	return c.service.Do()
}

type WordsController struct {
	*BaseController
}

type FeaturesController struct {
	*BaseController
}

func NewWordsController(service service.Service) *WordsController {
	return &WordsController{BaseController: &BaseController{service: service}}
}

func NewFeaturesController(service service.Service) *FeaturesController {
	return &FeaturesController{BaseController: &BaseController{service: service}}
}

func New(controllerType string, service service.Service) Controller {
	switch controllerType {
	case configuration.WordsLabel:
		return NewWordsController(service)
	case configuration.FeaturesLabel:
		return NewFeaturesController(service)
	default:
		return &BaseController{service: service}
	}
}

func (c *WordsController) GetAll() []string {
	return []string{"a", "b", "c"}
}
