package database

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Setup() error {
	//load variables

	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("Error loading Env file: %v", err)
	}
	mongoUrRL := os.Getenv("MONGO_URL")
	if mongoUrRL == "" {
		return fmt.Errorf("MONGO_URL is not set")
	}

	clientOptions := options.Client().ApplyURI(mongoUrRL)
	var err error
	client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return fmt.Errorf("Error conecting to the data base: %v", err)
	}
	fmt.Println("Connection to database established successfully")

	return nil
}

func GetClient() *mongo.Client {
	return client
}
