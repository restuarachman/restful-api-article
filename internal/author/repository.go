package author

import "tech-test/internal/author/entity"

type AuthorRepository interface {
	store(entity.Author) error
}
