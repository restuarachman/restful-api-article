package article

import (
	"tech-test/internal/article/dto"
)

type ArticleUsecase interface {
	Store(dto.ArticleRequest) error
	GetArticles() ([]dto.ArticleResponse, error)
}
