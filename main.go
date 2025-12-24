package main

import (
	"errors"
	"log"
	"net/url"
	"strings"
)

var ErrInvalidURI = errors.New("invalid URI")

func validateURL(rawURL string) (string, error) {

	rawURL = strings.TrimSpace(rawURL)

	if rawURL == "" {
		return "", ErrInvalidURI
	}

	if !strings.Contains(rawURL, "://") {
		rawURL = "https://" + rawURL
	}

	parsedURL, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return "", err
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return "", ErrInvalidURI
	}

	if parsedURL.Host == "" {
		return "", ErrInvalidURI
	}

	return parsedURL.String(), nil
}

func main() {
	urls := []string{
		"example.com",
		"https://google.com",
		"ftp://example.com",
		"javascript:alert(1)",
		"",
	}

	for _, u := range urls {
		normalized, err := validateURL(u)
		if err != nil {
			log.Printf("INVALID: %s", u)
		} else {
			log.Printf("VALID  : %s -> %s", u, normalized)
		}
	}
}
