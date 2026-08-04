package service

import (
	"strconv"

	"url-shortener/internal/models"
	"url-shortener/internal/storage"
)

var counter = 1

func GenerateShortCode() string {
	code := "url" + strconv.Itoa(counter)
	counter++
	return code
}

func CreateShortURL(originalURL string) (models.URL, error) {

	shortCode := GenerateShortCode()

	url := models.URL{
		OriginalURL: originalURL,
		ShortCode:   shortCode,
	}

	storage.Save(url)

	return url, nil
}