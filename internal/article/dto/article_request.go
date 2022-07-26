package dto

type ArticleRequest struct {
	Author_ID uint   `json:"author_id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
}
