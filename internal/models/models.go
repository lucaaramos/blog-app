package models

import (
	"time"

	"github.com/google/uuid"
)

// Post representa una publicación en el blog
type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPost(title, content, author string) *Post {
	return &Post{
		ID:        uuid.New().String(), // Genera un nuevo UUID como ID
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
