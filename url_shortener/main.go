package main

import (
	"log"
	"net/http"
	"os"

	"github.com/pandeyyyy/url_shortener/database"
	"github.com/pandeyyyy/url_shortener/handlers"
)

func main() {
	database.InitDB()
	defer database.CloseDB()
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	server := http.Server{
		Addr:    ":" + port,
		Handler: http.DefaultServeMux,
	}
	http.HandleFunc("/shorten", handlers.Shorten)
	http.HandleFunc("/", handlers.Resolve)

	log.Printf("Starting server on port %s\n", port)
	log.Fatal(server.ListenAndServe())
}
