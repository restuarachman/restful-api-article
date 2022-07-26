package usecase

import (
	"errors"
	"tech-test/internal/article"
	"tech-test/internal/article/dto"
	"tech-test/internal/article/entity"
	"time"
)

type ArticleUsecaseImpl struct {
	ar article.ArticleRepository
}

func NewArticleUsecase(ar article.ArticleRepository) article.ArticleUsecase {
	return &ArticleUsecaseImpl{ar}
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

func (au *ArticleUsecaseImpl) GetArticles() ([]dto.ArticleResponse, error) {
	articles, err := au.ar.Get()
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
