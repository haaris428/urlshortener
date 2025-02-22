package main

import (
	"net/http"
	"urlshortener/handlers"
)

func main() {
	// Setup routes
	http.HandleFunc("/shorten", handlers.ShortenHandler)
	http.HandleFunc("/", handlers.RedirectHandler)

	// Start periodic cleanup (optional)
	go handlers.CleanupExpiredURLs()

	// Run the HTTP server
	http.ListenAndServe(":8080", nil)
}
