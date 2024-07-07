package service

type Service interface {
	Do() string
}

type BaseService struct {
}
