package article

import (
	"tech-test/internal/article/entity"
)

type ArticleRepository interface {
	Store(entity.Article) error
	Get() ([]entity.Article, error)
}
