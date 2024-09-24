package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"todo/cmd/api/config"
	"todo/cmd/api/routes"
)

func main() {

	staticDir, err := filepath.Abs("./static")
	fmt.Println(staticDir)
	if err != nil {
		log.Fatal(err)
	}

	// Create a file server handler
	fs := http.FileServer(http.Dir(staticDir))

	// Handle requests to /static/
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	config.ConnectDB()

	router := routes.RegisterRoutes()

	log.Fatal(http.ListenAndServe(":8000", router))
}
