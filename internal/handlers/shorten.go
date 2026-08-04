package handlers

import (
	"encoding/json"
	"net/http"

	"url-shortener/internal/models"
	"url-shortener/internal/service"
	"url-shortener/internal/storage"
)

type ShortenRequest struct {
	URL string `json:"url"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {

	// Only allow POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	// Create a request object for THIS request only
	var req ShortenRequest

	// Decode JSON
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Validate
	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	// Generate short code
	shortCode := service.GenerateShortCode()

	// Create our model
	url := models.URL{
		OriginalURL: req.URL,
		ShortCode:   shortCode,
	}

	// Save it
	storage.Save(url)

	// Return response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(url)
}
