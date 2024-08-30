package models

import (
	"time"

	"github.com/google/uuid"
)

// Post representa una publicación en el blog
type Post struct {
	ID        string    `json:"id" bson:"_id"`
	Title     string    `json:"title" bson:"title"`
	Content   string    `json:"content" bson:"content"`
	Author    string    `json:"author" bson:"author"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
	Status    bool      `json:"status" bson:"status"`
	UserId    string    `json:"user_id" bson:"user_id"`
}

func NewPost(title, content, author, userid string) *Post {
	return &Post{
		ID:        uuid.New().String(), // Genera un nuevo UUID como ID
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    true,
		UserId:    userid,
	}
}
