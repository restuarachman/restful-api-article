package repository

import (
	"database/sql"
	"tech-test/internal/article"
	"tech-test/internal/article/entity"
)

type ArticleRepositoryImpl struct {
	db *sql.DB
}

func NewArticleRepo(db *sql.DB) article.ArticleRepository {
	return &ArticleRepositoryImpl{db}
}

func (ar *ArticleRepositoryImpl) Store(article entity.Article) error {
	insert, err := ar.db.Query("INSERT INTO articles (author_id, title, body, created_at) VALUES (?,?,?,?)", article.Author_ID, article.Title, article.Body, article.Created_at)

	if err != nil {
		return err
	}
	defer insert.Close()
	return nil

}

func (ar *ArticleRepositoryImpl) Get(query string, authorId, offset, limit int) ([]entity.Article, error) {
	var results *sql.Rows
	var err error
	if authorId != 0 {
		results, err = ar.db.Query("SELECT * FROM articles AS a WHERE (a.body LIKE ? OR a.title LIKE ?) AND author_id = ? ORDER BY created_at DESC LIMIT ?, ?", "%"+query+"%", "%"+query+"%", authorId, offset, limit)
		if err != nil {
			return []entity.Article{}, err
		}
	} else {
		results, err = ar.db.Query("SELECT * FROM articles AS a WHERE a.body LIKE ? OR a.title LIKE ? ORDER BY created_at DESC LIMIT ?, ?", "%"+query+"%", "%"+query+"%", offset, limit)
		if err != nil {
			return []entity.Article{}, err
		}
	}

	articles := []entity.Article{}
	for results.Next() {
		var articleRes entity.Article
		err = results.Scan(&articleRes.ID, &articleRes.Author_ID, &articleRes.Title, &articleRes.Body, &articleRes.Created_at)
		if err != nil {
			return []entity.Article{}, err
		}
		articles = append(articles, articleRes)
	}

	return articles, nil
}
