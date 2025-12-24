// Package service contains the business logic for the URL shortener.
package service

import (
	"crypto/rand"
	"encoding/hex"
	"mini-bitly/internal/domain"
	"mini-bitly/internal/repository"
	"net/url"
	"strings"
)

// ShortenerService handles the core business logic of shortening and resolving URLs.
// It acts as an intermediary between the HTTP handler and the repository.
type ShortenerService struct {
	repo repository.Repository
}

// NewShortenerService creates a new instance of ShortenerService with the given repository.
func NewShortenerService(repo repository.Repository) *ShortenerService {
	return &ShortenerService{repo: repo}
}

// Shorten validates the input URL, generates a unique code, and stores the mapping.
// It normalizes the URL (e.g., adds https://) and ensures strict validation (Host requirement).
func (s *ShortenerService) Shorten(rawURL string) (string, error) {
	// Validation Logic
	rawURL = strings.TrimSpace(rawURL)
	if rawURL == "" {
		return "", domain.ErrInvalidURL
	}
	// Default to HTTPS if no scheme is provided
	if !strings.Contains(rawURL, "://") {
		rawURL = "https://" + rawURL
	}
	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return "", domain.ErrInvalidURL
	}
	// Only allow HTTP/S schemes
	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return "", domain.ErrInvalidURL
	}
	// Ensure a host is present (e.g., prevent "https://")
	if parsedURL.Host == "" {
		return "", domain.ErrInvalidURL
	}

	// Generate Short Code
	code, err := generateRandomCode(6)
	if err != nil {
		return "", err
	}

	link := &domain.Link{
		Code:        code,
		OriginalURL: parsedURL.String(),
	}

	if err := s.repo.Save(link); err != nil {
		return "", err
	}

	return code, nil
}

// Resolve looks up the original URL associated with the given short code.
// Returns an error if the code is invalid or not found.
func (s *ShortenerService) Resolve(code string) (string, error) {
	link, err := s.repo.FindByCode(code)
	if err != nil {
		return "", err
	}
	return link.OriginalURL, nil
}

// generateRandomCode generates a crypto-random hex string of the specified length.
// Note: Length must be even as hex encoding doubles the characters.
func generateRandomCode(length int) (string, error) {
	bytes := make([]byte, length/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
