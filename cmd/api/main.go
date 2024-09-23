package main

import (
	"log"
	"net/http"
	"todo/cmd/api/config"
	"todo/cmd/api/routes"
)

func main() {
	// Initialize the database connection
	config.ConnectDB()

	// Register routes
	router := routes.RegisterRoutes()

	// Start the server
	log.Fatal(http.ListenAndServe(":8000", router))
}
