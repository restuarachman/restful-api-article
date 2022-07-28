package article

import (
	"tech-test/internal/article/entity"
)

type ArticleRepository interface {
	Store(entity.Article) error
	Get(string, int, int, int) ([]entity.Article, error)
}
