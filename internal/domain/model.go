// Package domain contains the core business models and errors.
// It is the heart of the application and has no external dependencies.
package domain

import "errors"

var (
	// ErrNotFound is returned when a requested link code does not exist in the repository.
	ErrNotFound = errors.New("link not found")

	// ErrInvalidURL is returned when the provided URL is malformed or empty.
	ErrInvalidURL = errors.New("invalid URL format")
)

// Link represents a shortened URL.
// It maps a unique short Code to an OriginalURL.
type Link struct {
	// Code is the unique identifier for the shortened link (e.g., "a1b2c3").
	Code string `json:"code"`

	// OriginalURL is the full URL that the code redirects to (e.g., "https://google.com").
	OriginalURL string `json:"original_url"`
}
