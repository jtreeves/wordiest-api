package service

import "github.com/jtreeves/wordiest-api/pkg/model"

type Service interface {
	Do() model.Response
}

type BaseService struct {
}
