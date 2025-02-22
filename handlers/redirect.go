package handlers

import (
	"net/http"
	"time"
	"urlshortener/models"
)

func RedirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:] // Get the short URL

	models.Store.Mutex.RLock()
	urlData, exists := models.Store.Mapping[shortURL]
	models.Store.Mutex.RUnlock()

	if !exists || time.Now().After(urlData.ExpiresAt) {
		http.Error(w, "URL not found or expired", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, urlData.LongURL, http.StatusFound)
}
