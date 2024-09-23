package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TodoCollection *mongo.Collection

func ConnectDB() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@adw8.qwutgju.mongodb.net/?retryWrites=true&w=majority&appName=adw8", dbUsername, dbPassword)

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	TodoCollection = client.Database("lf").Collection("todo")
	fmt.Println("Connected to MongoDB!")
}
