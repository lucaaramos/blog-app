package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post representa una publicación en el blog
type Post struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Title     string             `json:"title" bson:"title"`
	Content   string             `json:"content" bson:"content"`
	Author    string             `json:"author" bson:"author"`
	CreatedAt time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time          `json:"updated_at" bson:"updated_at"`
	Status    bool               `json:"status" bson:"status"`
	UserId    string             `json:"user_id" bson:"user_id"`
}

func NewPost(title, content, author, userid string) *Post {
	return &Post{
		Title:     title,
		Content:   content,
		Author:    author,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Status:    true,
		UserId:    userid,
	}
}
