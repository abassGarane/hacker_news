package models

import "time"

type NewsItem struct {
	ID        int64      `json:"_id"`
	NewsLink  string     `json:"news_link"`
	Comments  []*Comment `json:"comments"`
	Likes     int64      `json:"likes"`
	Dislikes  int64      `json:"dislikes"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
