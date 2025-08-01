package helpers

import (
	"os"
	"strings"
)

// Takes a URL string and adds "http://" if it's missing.
func EnforceHTTP(url string) string {
	if url[:4] != "http" {
		return "http://" + url
	}
	return url
}

// Checks if the provided URL matches the domain or is valid
func RemoveDomainError(url string) bool {
	domain := os.Getenv("DOMAIN") // e.g., "localhost:3000"

	if url == domain {
		return false
	}

	// Clean the URL by removing protocol and "www."
	newURL := strings.Replace(url, "http://", "", 1)
	newURL = strings.Replace(newURL, "https://", "", 1)
	newURL = strings.Replace(newURL, "www.", "", 1)
	newURL = strings.Split(newURL, "/")[0] // isolate domain

	// If cleaned domain matches your app domain, reject it
	return newURL != domain
}
