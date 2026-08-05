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

	err := storage.Save(url)
	if err != nil {
		return models.URL{}, err
	}

	return url, nil
}

func GetOriginalURL(shortCode string) (models.URL, error) {
	return storage.Get(shortCode)
}
