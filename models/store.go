package models

import (
	"sync"
	"time"
)

// URLData represents the long URL and its expiration time
type URLData struct {
	LongURL   string
	ExpiresAt time.Time
}

// URLStore represents an in-memory URL models
type URLStore struct {
	Mapping map[string]URLData
	Mutex   sync.RWMutex
}

// Global in-memory models (can be replaced with DB later)
var Store = URLStore{
	Mapping: make(map[string]URLData),
	Mutex:   sync.RWMutex{},
}
