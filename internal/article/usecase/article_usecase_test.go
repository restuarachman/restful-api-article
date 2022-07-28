package usecase

import (
	"errors"
	_articleDTO "tech-test/internal/article/dto"
	_articleEntity "tech-test/internal/article/entity"
	_mockArticleRepo "tech-test/internal/article/mocks"
	_authorEntity "tech-test/internal/author/entity"
	_mockAuthorRepo "tech-test/internal/author/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	mockArticleEntity = _articleEntity.Article{
		ID:         1,
		Author_ID:  1,
		Title:      "title",
		Body:       "body",
		Created_at: time.Now(),
	}

	mockArticleEntityList = []_articleEntity.Article{mockArticleEntity, mockArticleEntity, mockArticleEntity}

	mockArticleDTO = _articleDTO.ArticleRequest{
		Author_ID: 1,
		Title:     "title",
		Body:      "body",
	}

	mockAuthorEntity = _authorEntity.Author{
		ID:   1,
		Name: "dummy",
	}
)

func TestStore(t *testing.T) {
	mockArticleRepo := _mockArticleRepo.NewArticleRepository(t)

	t.Run("success", func(t *testing.T) {
		expectedArticleEntity := _articleEntity.Article{
			Author_ID:  1,
			Title:      "title",
			Body:       "body",
			Created_at: time.Now(),
		}
		mockArticleRepo.On("Store", expectedArticleEntity).Return(nil).Once()

		testArticleUsecase := NewArticleUsecase(mockArticleRepo, nil)

		err := testArticleUsecase.Store(mockArticleDTO)
		assert.NoError(t, err)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockArticleRepo.On("Store", mock.Anything).Return(errors.New("internal server error")).Once()

		testArticleUsecase := NewArticleUsecase(mockArticleRepo, nil)

		err := testArticleUsecase.Store(mockArticleDTO)
		assert.Error(t, err)
	})

}

func TestGetArticles(t *testing.T) {
	mockArticleRepo := _mockArticleRepo.NewArticleRepository(t)
	mockAuthorRepo := _mockAuthorRepo.NewAuthorRepository(t)

	t.Run("success", func(t *testing.T) {
		mockArticleRepo.On("Get", "test", 0, 1, 1).Return(mockArticleEntityList, nil).Once()

		testArticleUsecase := NewArticleUsecase(mockArticleRepo, nil)

		res, err := testArticleUsecase.GetArticles("test", "", 1, 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("success", func(t *testing.T) {
		mockAuthorRepo.On("Get", "dummy").Return(mockAuthorEntity, nil).Once()
		mockArticleRepo.On("Get", "test", 1, 1, 1).Return(mockArticleEntityList, nil).Once()

		testArticleUsecase := NewArticleUsecase(mockArticleRepo, mockAuthorRepo)

		res, err := testArticleUsecase.GetArticles("test", "dummy", 1, 1)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockAuthorRepo.On("Get", "dummy").Return(mockAuthorEntity, nil).Once()
		mockArticleRepo.On("Get", "test", 1, 1, 1).Return([]_articleEntity.Article{}, errors.New("internal server error")).Once()

		testArticleUsecase := NewArticleUsecase(mockArticleRepo, mockAuthorRepo)

		res, err := testArticleUsecase.GetArticles("test", "dummy", 1, 1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("internal server error", func(t *testing.T) {
		mockAuthorRepo.On("Get", "dummy").Return(_authorEntity.Author{}, errors.New("internal server error")).Once()

		testArticleUsecase := NewArticleUsecase(mockArticleRepo, mockAuthorRepo)

		res, err := testArticleUsecase.GetArticles("test", "dummy", 1, 1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

	t.Run("author not found", func(t *testing.T) {
		mockAuthorRepo.On("Get", "dummy").Return(_authorEntity.Author{}, nil).Once()

		testArticleUsecase := NewArticleUsecase(mockArticleRepo, mockAuthorRepo)

		res, err := testArticleUsecase.GetArticles("test", "dummy", 1, 1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}
