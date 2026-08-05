package main

import (
	"log"
	"net/http"

	"url-shortener/internal/handlers"
	"url-shortener/internal/storage"
)

func main() {
    storage.Connect()

    http.HandleFunc("/shorten", handlers.ShortenHandler)
    http.HandleFunc("/", handlers.RedirectHandler)

    log.Fatal(http.ListenAndServe(":8080", nil))
}