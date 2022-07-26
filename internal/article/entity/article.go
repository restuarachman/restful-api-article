package entity

import "time"

type Article struct {
	ID         uint
	Author_ID  uint
	Title      string
	Body       string
	Created_at time.Time
}
