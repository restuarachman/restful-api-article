package http

import (
	"net/http"
	"strconv"
	"tech-test/internal/article"
	"tech-test/internal/article/dto"

	"github.com/labstack/echo/v4"
)

type ArticleHandler struct {
	au article.ArticleUsecase
}

func NewArticleHandler(e *echo.Echo, au article.ArticleUsecase) {
	articleHandler := ArticleHandler{
		au: au,
	}

	e.POST("/articles", articleHandler.postArticle)
	e.GET("/articles", articleHandler.getArticle)
}

func (ah *ArticleHandler) postArticle(c echo.Context) error {
	articleReq := dto.ArticleRequest{}
	c.Bind(&articleReq)

	err := ah.au.Store(articleReq)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "OK",
	})
}

func (ah *ArticleHandler) getArticle(c echo.Context) error {
	query := c.QueryParam("query")
	author := c.QueryParam("author")

	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "bad param input",
		})
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "bad param input",
		})
	}

	articlesRes, err := ah.au.GetArticles(query, author, offset, limit)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":  "OK",
		"articles": articlesRes,
	})
}
