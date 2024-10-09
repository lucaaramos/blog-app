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
	post.ID = primitive.NewObjectID()
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()
	post.Status = true
	_, err := r.collection.InsertOne(ctx, post)
	return err
}

func (r *PostRepository) GetPostByID(ctx context.Context, id string) (*models.Post, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid post ID format")
	}

	var post models.Post
	if err := r.collection.FindOne(ctx, bson.M{"_id": objID, "status": true}).Decode(&post); err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("post not found")
		}
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) GetAllBlogs(ctx context.Context) ([]models.Post, error) {
	var posts []models.Post
	filter := bson.M{"status": true}
	cursor, err := r.collection.Find(ctx, filter)
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

func (r *PostRepository) UpdateBlog(ctx context.Context, id string, updatePost *models.Post) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid post ID format")
	}

	filter := bson.M{"_id": objID, "status": true}
	update := bson.M{
		"$set": bson.M{
			"title":   updatePost.Title,
			"content": updatePost.Content,
			"author":  updatePost.Author,
		},
	}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("Post not found")
	}
	return nil
}

// func (r *PostRepository) UpdateBlog(ctx context.Context, id string, updatePost *models.Post) error {
// 	// Intentar convertir el ID a un ObjectID
// 	objID, err := primitive.ObjectIDFromHex(id)
// 	if err != nil {
// 		return fmt.Errorf("invalid post ID format: %w", err)
// 	}

// 	// Crear el filtro para encontrar el post por su ID y que esté activo (status=true)
// 	filter := bson.M{"_id": objID, "status": true}

// 	// Actualizar el post con los nuevos valores y actualizar la fecha 'updatedAt'
// 	update := bson.M{
// 		"$set": bson.M{
// 			"title":     updatePost.Title,
// 			"content":   updatePost.Content,
// 			"author":    updatePost.Author,
// 			"updatedAt": time.Now(), // Actualizar el campo updatedAt
// 		},
// 	}

// 	// Intentar actualizar el documento
// 	result, err := r.collection.UpdateOne(ctx, filter, update)
// 	if err != nil {
// 		return fmt.Errorf("error updating post: %w", err)
// 	}

// 	// Si no se encontró ningún documento que coincida con el filtro, lanzar error
// 	if result.MatchedCount == 0 {
// 		return errors.New("post not found or already deleted")
// 	}

// 	return nil
// }

func (r *PostRepository) DeleteBlog(ctx context.Context, id string, deletePost *models.Post) error {
	filter := bson.M{"id": id, "status": true}
	update := bson.M{
		"$set": bson.M{
			"status": deletePost.Status,
		},
	}
	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if result.MatchedCount == 0 {
		return errors.New("Post not found")
	}
	return nil
}
