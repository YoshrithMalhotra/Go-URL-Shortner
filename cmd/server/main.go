package main

import (
	"log"
	"net/http"

	"url-shortener/internal/handlers"
)

func main() {

	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.RedirectHandler)

	log.Println("🚀 Server running on http://localhost:8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}