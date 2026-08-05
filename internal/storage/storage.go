package storage

import "url-shortener/internal/models"


func Save(url models.URL) error {
	_, err := DB.Exec("INSERT INTO urls (original_url, short_code) VALUES ($1, $2)", url.OriginalURL, url.ShortCode)
	if err != nil {
		return err
	}
	return nil
}

func Get(shortCode string) (models.URL, error) {
	var url models.URL
	row := DB.QueryRow("SELECT original_url, short_code FROM urls WHERE short_code = $1", shortCode)
	err := row.Scan(&url.OriginalURL, &url.ShortCode)
	if err != nil {
		return models.URL{}, err
	}
	return url, nil
}
