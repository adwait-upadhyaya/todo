package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"todo/cmd/api/config"
	"todo/cmd/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	cursor, err := config.TodoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		http.Error(w, "Error fetching todos", http.StatusInternalServerError)
		return
	}
	var todos []models.Todo
	if err = cursor.All(context.Background(), &todos); err != nil {
		http.Error(w, "Error decoding todos", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	if newTodo.Title == "" {
		http.Error(w, "Missing required field: title", http.StatusBadRequest)
		return
	}
	newTodo.ID = primitive.NewObjectID()
	result, err := config.TodoCollection.InsertOne(context.TODO(), newTodo)
	if err != nil {
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
		return
	}
	fmt.Println(result)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newTodo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	update := bson.D{{"$set", bson.D{{"completed", true}}}}
	filter := bson.D{{"_id", objectId}}

	result, err := config.TodoCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w, "Error updating todo", http.StatusInternalServerError)
		return
	}
	fmt.Println(result)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo updated successfully"})
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")

	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	result, err := config.TodoCollection.DeleteOne(context.TODO(), bson.M{"_id": objectId})
	if err != nil {
		http.Error(w, "Error deleting todo", http.StatusInternalServerError)
		return
	}

	if result.DeletedCount == 0 {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted successfully"})
}
