package models

type Comment struct {
	ID       int64      `json:"id"`
	Body     string     `json:"body"`
	Likes    int64      `json:"likes"`
	Dislikes int64      `json:"dislikes"`
	Comments []*Comment `json:"comments"`
}
