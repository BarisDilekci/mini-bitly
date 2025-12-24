// Package main is the entry point for the mini-bitly URL shortener service.
// It orchestrates the startup, dependency injection, and server lifeycle.
package main

import (
	"log"
	"mini-bitly/internal/handler"
	"mini-bitly/internal/repository"
	"mini-bitly/internal/service"
	"net/http"
)

func main() {
	// Initialize the repository layer (In-Memory)
	repo := repository.NewMemoryRepository()

	// Initialize the business logic layer with the repository
	svc := service.NewShortenerService(repo)

	// Initialize the HTTP handler with the service
	h := handler.NewHandler(svc)

	// Setup HTTP routing
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", h.Shorten)
	mux.HandleFunc("/", h.Redirect)

	// Start the server
	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
