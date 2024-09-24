package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"todo/cmd/api/config"
	"todo/cmd/api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func RenderTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := GetTodos()
	if err != nil {
		http.Error(w, "Error fetching todos", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("cmd/api/templates/index.gohtml"))

	tmpl.Execute(w, struct {
		Todos []models.Todo
	}{
		Todos: todos,
	})
}

func GetTodos() ([]models.Todo, error) {
	cursor, err := config.TodoCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	var todos []models.Todo
	if err = cursor.All(context.Background(), &todos); err != nil {
		return nil, err
	}

	return todos, nil
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, "Invalid form data", http.StatusBadRequest)
		return
	}

	newTodo.ID = primitive.NewObjectID()

	if newTodo.Title == "" {
		http.Error(w, "Missing required field: title", http.StatusBadRequest)
		return
	}

	result, err := config.TodoCollection.InsertOne(context.TODO(), newTodo)
	fmt.Println(result)
	if err != nil {
		http.Error(w, "Error creating todo", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	id := params.Get("id")
	response := struct {
		Message string
	}{
		Message: "Updated succesfully",
	}

	update := bson.D{{"$set", bson.D{{"completed", true}}}}

	if id != "" {
		objectId, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "Invalid ID format", http.StatusBadRequest)
			return
		}
		filter := bson.D{{"_id", objectId}}
		result, err := config.TodoCollection.UpdateOne(context.TODO(), filter, update)
		fmt.Println(result)
		if err != nil {
			http.Error(w, "Error Updating Todo", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	} else {
		// update all logic here
		filter := bson.D{}
		_, err := config.TodoCollection.UpdateMany(context.TODO(), filter, update)

		if err != nil {
			http.Error(w, "Error Updating todo", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}
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
