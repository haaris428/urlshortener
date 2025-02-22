package handlers

import (
	"encoding/json"
	"net/http"
	"time"
	"urlshortener/models"
	"urlshortener/utils"
)

type ShortenRequest struct {
	LongURL     string `json:"long_url"`
	CustomAlias string `json:"custom_alias,omitempty"`
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

func ShortenHandler(w http.ResponseWriter, r *http.Request) {
	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid Request Payload", http.StatusBadRequest)
		return
	}

	isRegisteredUser := r.Header.Get("Authorization") == "registered_user_token" // Replace with real authentication logic

	// Set expiration based on user type
	var expiration time.Time
	if isRegisteredUser {
		expiration = time.Now().Add(30 * 24 * time.Hour) // Registered user
	} else {
		expiration = time.Now().Add(24 * time.Hour) // Anonymous user
	}

	// Generate the alias (custom alias or random)
	shortURL := req.CustomAlias
	if shortURL != "" {
		// Check if custom alias is already taken
		models.Store.Mutex.RLock()
		_, exists := models.Store.Mapping[shortURL]
		models.Store.Mutex.RUnlock()
		if exists {
			http.Error(w, "Alias already taken. Please try another.", http.StatusConflict)
			return
		}
	} else {
		shortURL = utils.GenerateShortURL()
	}

	models.Store.Mutex.Lock()
	models.Store.Mapping[shortURL] = models.URLData{
		LongURL:   req.LongURL,
		ExpiresAt: expiration,
	}
	models.Store.Mutex.Unlock()

	response := ShortenResponse{ShortURL: "http://localhost:8080/" + shortURL}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
