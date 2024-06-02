package repository

import (
	"blog/internal/models"
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PostRepository struct {
	collection *mongo.Collection
}

func NewPostRepository() *PostRepository {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	collection := client.Database("blog").Collection("posts")
	return &PostRepository{collection}
}

func (r *PostRepository) CreatePost(ctx context.Context, post *models.Post) error {
	post.ID = primitive.NewObjectID().Hex()
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, post)
	return err
}

func (r *PostRepository) GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	var post models.Post
	if err := r.collection.FindOne(ctx, bson.M{"id": id}).Decode(&post); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &post, nil
}
