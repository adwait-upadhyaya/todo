package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type todo struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Title     string             `bson:"title"`
	Completed bool               `bson:"completed"`
}

var todoCollection *mongo.Collection

func connectDB() {
	envErr := godotenv.Load(".env")
	if envErr != nil {
		log.Fatalf("Error loading .env file: %s", envErr)
	}
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	connectionString := fmt.Sprintf("mongodb+srv://%v:%v@adw8.qwutgju.mongodb.net/?retryWrites=true&w=majority&appName=adw8", db_username, db_password)

	client, err := mongo.Connect(context.TODO(), options.Client().
		ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}

	// Check the connection.
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	} else {
		todoCollection = client.Database("lf").Collection("todo")
		fmt.Println("Connected to mongoDB!!!")
	}

}

func getTodos(w http.ResponseWriter, r *http.Request) {
	cursor, err := todoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, "Error Fetching Todos", http.StatusInternalServerError)
		return
	}

	var todos []todo

	if err = cursor.All(context.Background(), &todos); err != nil {
		panic(err)
	}

	json.NewEncoder(w).Encode(todos)
	w.WriteHeader(http.StatusOK)
}

func main() {
	connectDB()

	http.HandleFunc("/todos", getTodos)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
