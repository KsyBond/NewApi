package models

type Post struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Title       string `json:"title"`
}
