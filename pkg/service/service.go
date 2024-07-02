package service

type Service interface {
	Do() string
}

type BaseService struct {
}

func (s *BaseService) Do() string {
	return ""
}
