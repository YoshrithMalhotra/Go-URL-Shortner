package storage

import (
	"sync"

	"url-shortener/internal/models"
)

var (
	mu    sync.RWMutex
	items = make(map[string]models.URL)
)

func Save(url models.URL) {
	mu.Lock()
	defer mu.Unlock()

	items[url.ShortCode] = url
}

func Get(shortCode string) (models.URL, bool) {
	mu.RLock()
	defer mu.RUnlock()

	url, ok := items[shortCode]
	return url, ok
}
