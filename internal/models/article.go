package models

type Article struct {
	ID        int    `json:"id"`
	Title     string `json:"title" validate:"required"`
	Content   string `json:"content" validate:"required"`
	ImageURL  string `json:"image_url"`
	CreatedAt string `json:"created_at"`
}