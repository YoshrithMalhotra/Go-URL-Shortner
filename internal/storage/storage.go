package storage

import "url-shortener/internal/models"

var urls = make(map[string]models.URL)

func Save(url models.URL) {
	urls[url.ShortCode] = url
}

func Get(shortCode string) (models.URL, bool) {
	url, exists := urls[shortCode]
	return url, exists
}