package main

import (
	"log"
	"net/http"
	"todo/cmd/api/config"
	"todo/cmd/api/routes"
)

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	config.ConnectDB()

	router := routes.RegisterRoutes()

	log.Fatal(http.ListenAndServe(":8000", router))
}
