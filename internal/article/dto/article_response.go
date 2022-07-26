package dto

import "time"

type ArticleResponse struct {
	ID         uint      `json:"ID"`
	Author_ID  uint      `json:"author_id"`
	Title      string    `json:"title"`
	Body       string    `json:"body"`
	Created_at time.Time `json:"created_at"`
}
