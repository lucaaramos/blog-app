package repository

import (
	"blog/internal/models"
	"context"
	"errors"
	"log"
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

func (r *PostRepository) GetAllBlogs(ctx context.Context) ([]models.Post, error) {
	var posts []models.Post

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Println("Error getting data", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var post models.Post
		if err := cursor.Decode(&post); err != nil {
			log.Println("Error decoding data", err)
			continue
		}
		posts = append(posts, post)
	}

	if err := cursor.Err(); err != nil {
		log.Println("Cursor error: ", err)
		return nil, err
	}

	return posts, nil

}
