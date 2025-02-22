package handlers

import (
	"time"
	"urlshortener/models"
)

func CleanupExpiredURLs() {
	for {
		time.Sleep(1 * time.Hour) // Perform cleanup once every hour
		models.Store.Mutex.Lock()
		for key, value := range models.Store.Mapping {
			if time.Now().After(value.ExpiresAt) {
				delete(models.Store.Mapping, key) // Remove expired URLs
			}
		}
		models.Store.Mutex.Unlock()
	}
}
