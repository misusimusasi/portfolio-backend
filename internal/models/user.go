package models

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username" validate:"required,min=3"`
	PasswordHash string `json:"-" validate:"required,min=6"`
}