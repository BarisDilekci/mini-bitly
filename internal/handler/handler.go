// Package handler implements the HTTP transport layer.
package handler

import (
	"encoding/json"
	"errors"
	"mini-bitly/internal/domain"
	"mini-bitly/internal/service"
	"net/http"
	"strings"
)

// Handler manages HTTP requests and responses.
type Handler struct {
	svc *service.ShortenerService
}

// NewHandler creates a new Handler instance with the given service.
func NewHandler(svc *service.ShortenerService) *Handler {
	return &Handler{svc: svc}
}

// ShortenRequest defines the expected JSON body for shortening a URL.
type ShortenRequest struct {
	URL string `json:"url"`
}

// ShortenResponse defines the JSON response containing the shortened URL.
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// Shorten handles POST requests to create a new short URL.
// It expects a JSON body with a "url" field and returns the shortened link.
func (h *Handler) Shorten(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req ShortenRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	code, err := h.svc.Shorten(req.URL)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidURL) {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	resp := ShortenResponse{
		ShortURL: "http://" + r.Host + "/" + code,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// Redirect handles GET requests to redirect to the original URL.
// It extracts the code from the URL path and issues a 302 Found redirect.
func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	code := strings.TrimPrefix(r.URL.Path, "/")
	if code == "" {
		http.Error(w, "Missing code", http.StatusBadRequest)
		return
	}

	originalURL, err := h.svc.Resolve(code)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.NotFound(w, r)
			return
		}
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, originalURL, http.StatusFound)
}
