package repository

import (
	"database/sql"
	"fmt"
	"tech-test/internal/author"
	"tech-test/internal/author/entity"
)

type AuthorRepositoryImpl struct {
	db *sql.DB
}

func NewAuthorRepository(db *sql.DB) author.AuthorRepository {
	return &AuthorRepositoryImpl{db}
}

func (authorRepo *AuthorRepositoryImpl) Get(name string) (entity.Author, error) {
	author := entity.Author{}
	row := authorRepo.db.QueryRow("SELECT * FROM authors WHERE name = ?", name)
	if err := row.Scan(&author.ID, &author.Name); err != nil {
		if err == sql.ErrNoRows {
			return author, fmt.Errorf("authorName %s: no such author", name)
		}
		return author, fmt.Errorf("authorName %s: %v", name, err)
	}
	return author, nil
}
