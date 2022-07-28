package author

import "tech-test/internal/author/entity"

type AuthorRepository interface {
	Get(string) (entity.Author, error)
}
