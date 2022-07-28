package article

import (
	"tech-test/internal/article/dto"
)

type ArticleUsecase interface {
	Store(dto.ArticleRequest) error
	GetArticles(string, string, int, int) ([]dto.ArticleResponse, error)
}
