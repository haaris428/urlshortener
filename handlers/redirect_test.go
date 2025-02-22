package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"urlshortener/models"
)

func TestRedirectHandler(t *testing.T) {
	// Setup store with a test entry
	shortURL := "abc123"
	longURL := "https://example.com"
	models.Store.Mutex.Lock()
	models.Store.Mapping[shortURL] = models.URLData{
		LongURL:   longURL,
		ExpiresAt: time.Now().Add(1 * time.Hour),
	}
	models.Store.Mutex.Unlock()

	// Test the redirect handler
	req := httptest.NewRequest(http.MethodGet, "/"+shortURL, nil)
	w := httptest.NewRecorder()

	RedirectHandler(w, req)

	// Verify redirect
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusFound {
		t.Fatalf("expected status 302 Found, got %d", resp.StatusCode)
	}
	if resp.Header.Get("Location") != longURL {
		t.Errorf("expected redirect to %s, got %s", longURL, resp.Header.Get("Location"))
	}
}

func TestRedirectHandler_Expired(t *testing.T) {
	// Setup store with an expired entry
	shortURL := "expired123"
	models.Store.Mutex.Lock()
	models.Store.Mapping[shortURL] = models.URLData{
		LongURL:   "https://example.com",
		ExpiresAt: time.Now().Add(-1 * time.Hour), // Already expired
	}
	models.Store.Mutex.Unlock()

	// Test the redirect handler
	req := httptest.NewRequest(http.MethodGet, "/"+shortURL, nil)
	w := httptest.NewRecorder()

	RedirectHandler(w, req)

	// Verify 404 for expired URL
	resp := w.Result()
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Fatalf("expected status 404 Not Found for expired URL, got %d", resp.StatusCode)
	}
}
