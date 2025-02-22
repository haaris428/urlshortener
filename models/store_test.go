package models

import (
	"testing"
	"time"
)

func TestURLStore(t *testing.T) {
	// Create a test store
	store := URLStore{
		Mapping: make(map[string]URLData),
	}
	
	// Test adding an entry
	shortURL := "abc123"
	longURL := "https://example.com"
	expiration := time.Now().Add(24 * time.Hour)
	store.Mutex.Lock()
	store.Mapping[shortURL] = URLData{LongURL: longURL, ExpiresAt: expiration}
	store.Mutex.Unlock()

	store.Mutex.RLock()
	data, exists := store.Mapping[shortURL]
	store.Mutex.RUnlock()

	if !exists {
		t.Fatalf("expected URL to exist in store")
	}
	if data.LongURL != longURL {
		t.Errorf("expected long URL '%s', got '%s'", longURL, data.LongURL)
	}
	if data.ExpiresAt.Before(time.Now()) {
		t.Errorf("expected expiration time to be in the future")
	}
}