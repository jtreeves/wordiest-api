package service

import "github.com/jtreeves/wordiest-api/pkg/model"

type WordService struct {
	*BaseService
}

func NewWordService() *WordService {
	return &WordService{}
}

func (s *WordService) Do() model.Response {
	w := model.Word{
		ID:     1,
		Value:  "word",
		Length: 4,
		Letter: "W",
	}
	r := model.Response{Data: w}

	return r
}
