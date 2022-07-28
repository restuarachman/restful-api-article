package usecase

import (
	"errors"
	"tech-test/internal/article"
	"tech-test/internal/article/dto"
	"tech-test/internal/article/entity"
	"tech-test/internal/author"
	entityAuthor "tech-test/internal/author/entity"
	"time"
)

type ArticleUsecaseImpl struct {
	ar         article.ArticleRepository
	authorRepo author.AuthorRepository
}

func NewArticleUsecase(ar article.ArticleRepository, authorRepo author.AuthorRepository) article.ArticleUsecase {
	return &ArticleUsecaseImpl{ar, authorRepo}
}

func (au *ArticleUsecaseImpl) Store(article dto.ArticleRequest) error {

	articleEntity := entity.Article{
		Author_ID:  article.Author_ID,
		Title:      article.Title,
		Body:       article.Body,
		Created_at: time.Now(),
	}

	err := au.ar.Store(articleEntity)
	if err != nil {
		return errors.New("internal server error")
	}

	return nil
}

func (au *ArticleUsecaseImpl) GetArticles(query, author string, offset, limit int) ([]dto.ArticleResponse, error) {
	var authorEntity entityAuthor.Author
	var err error

	if author != "" {
		authorEntity, err = au.authorRepo.Get(author)
		if err != nil {
			return []dto.ArticleResponse{}, err
		}
		if authorEntity.ID == 0 {
			return []dto.ArticleResponse{}, errors.New("author not found")
		}
	}

	articles, err := au.ar.Get(query, int(authorEntity.ID), offset, limit)
	if err != nil {
		return []dto.ArticleResponse{}, err
	}

	var articlesRes []dto.ArticleResponse

	for _, val := range articles {
		articlesRes = append(articlesRes, dto.ArticleResponse{
			ID:         val.ID,
			Author_ID:  val.Author_ID,
			Title:      val.Title,
			Body:       val.Body,
			Created_at: val.Created_at,
		})
	}

	return articlesRes, nil
}
