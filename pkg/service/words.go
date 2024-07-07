package service

type WordsService struct {
	*BaseService
}

func NewWordsService() *WordsService {
	return &WordsService{}
}

func (s *WordsService) Do() string {
	return "words"
}
