package models

import "time"

// Post representa una publicación en el blog
type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Author    string    `json:"author"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewPost crea una nueva publicación con los datos proporcionados
func NewPost(title, content, author string) *Post {
	return &Post{
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}
