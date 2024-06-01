package repository

import (
	"blog/internal/models"
	"context"

	// "errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type PostRepository struct {
	collection *mongo.Collection
}

func NewPostRepository(collection *mongo.Collection) *PostRepository {
	return &PostRepository{}
}

func (r *PostRepository) CreatePost(ctx context.Context, post models.Post) error {
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	_, err := r.collection.InsertOne(ctx, post)
	if err != nil {
		return err
	}
	return nil
}

func (r *PostRepository) GetId(ctx context.Context, get models.Post) error {
	filter := bson.M{"_id": get.ID}
	err := r.collection.FindOne(ctx, filter).Decode(&get)
	if err != nil {
		return err
	}
	return nil
}
