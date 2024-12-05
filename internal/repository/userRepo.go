package repository

import (
	"blog/internal/models"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepository struct {
	collection     *mongo.Collection
	postCollection *mongo.Collection
}

func NewUserRepository() *UserRepository {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		panic(err)
	}

	userCollection := client.Database("blog").Collection("users")
	postCollection := client.Database("blog").Collection("posts")

	return &UserRepository{
		collection:     userCollection,
		postCollection: postCollection,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) error {
	user.Id = primitive.NewObjectID().Hex()
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	// hasheamos la contraseña antes de guardar

	if err := user.HashPassword(); err != nil {
		return err
	}

	// Insertar el usuario en la colección "users"
	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Error al insertar usuario en la base de datos: %v\n", err)
		return err
	}

	return nil
}

// creamos una funcion para buscar el usuario por username
func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	filter := bson.M{"username": username}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	var users []models.User
	filter := bson.M{}
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		log.Print("Error getting users", err)
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user models.User
		if err := cursor.Decode(&user); err != nil {
			log.Print("Error decoding data", err)
			continue
		}
		users = append(users, user)
	}
	if err := cursor.Err(); err != nil {
		log.Print("Error iterating over cursor", err)
		return nil, err
	}
	return users, nil
}
