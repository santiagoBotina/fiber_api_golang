package db

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

var DB *mongo.Database

func Initialize() {
	err := godotenv.Load()
	if err != nil {
		return
	}

	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_NAME := os.Getenv("DB_NAME")

	URL := fmt.Sprintf(
		"mongodb://%s:%s",
		DB_HOST, DB_PORT,
	)

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URL))

	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Database connected")

	DB = client.Database(DB_NAME)
}
