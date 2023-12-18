package models

import (
	"time"

	"github.com/uptrace/bun"
)

type NewsItem struct {
	bun.BaseModel `bun:"table:news_items,alias:ni"`
	ID            int64     `json:"_id" bun:"id,pk,autoincrement"`
	NewsLink      string    `json:"news_link" bun:"news_link,notnull"`
	Likes         int64     `json:"likes" bun:"likes"`
	Dislikes      int64     `json:"dislikes" bun:"dislikes"`
	CreatedAt     time.Time `json:"created_at" bun:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bun:"updated_at"`
}
