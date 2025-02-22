package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"urlshortener/models"
)

func TestShortenHandler_AnonymousUser(t *testing.T) {
	// Reset the store for test isolation
	models.Store = models.URLStore{
		Mapping: make(map[string]models.URLData),
	}

	// Prepare the test request payload
	requestBody := ShortenRequest{
		LongURL: "https://example.com",
	}
	body, _ := json.Marshal(requestBody)

	req := httptest.NewRequest(http.MethodPost, "/shorten", bytes.NewReader(body))
	w := httptest.NewRecorder()

	ShortenHandler(w, req)

	// Verify the response
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var response ShortenResponse
	json.NewDecoder(resp.Body).Decode(&response)
	if response.ShortURL == "" {
		t.Errorf("expected a valid short URL in response")
	}

	// Verify the store was updated
	models.Store.Mutex.RLock()
	defer models.Store.Mutex.RUnlock()

	found := false
	for _, urlData := range models.Store.Mapping {
		if urlData.LongURL == requestBody.LongURL {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("expected data in the store but it's missing")
	}
}
