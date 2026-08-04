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

	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if req.URL == "" {
		http.Error(w, "URL is required", http.StatusBadRequest)
		return
	}

	shortCode := service.GenerateShortCode()

	url := models.URL{
		OriginalURL: req.URL,
		ShortCode:   shortCode,
	}

	storage.Save(url)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(url)
}
