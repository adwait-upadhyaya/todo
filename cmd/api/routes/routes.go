package routes

import (
	"net/http"
	"todo/cmd/api/handlers"
)

func RegisterRoutes() *http.ServeMux {

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.RenderTodos)
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			handlers.GetTodos()
		case http.MethodPost:
			handlers.CreateTodo(w, r)
		case http.MethodPut:
			handlers.UpdateTodo(w, r)
		case http.MethodDelete:
			handlers.DeleteTodo(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	return mux
}
