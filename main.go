package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// Handler for the root route
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Welcome to the URL Shortener Service!"))
}

func main() {
	// Setup environment variables, if any (e.g., for PORT)
	port := getPort()

	// Setup routes
	http.HandleFunc("/", homeHandler)

	// Start the server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      nil,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server is starting on port %s...", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

// getPort retrieves the port from an environment variable or defaults to 8080
func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return port
}
