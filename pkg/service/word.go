package service

type WordService struct {
	*BaseService
}

func NewWordService() *WordService {
	return &WordService{}
}

func (s *WordService) Do() string {
	return "word"
}
