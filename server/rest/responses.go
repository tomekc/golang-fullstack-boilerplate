// Package rest provides helpers for building HTTP responses and serializing them to JSON.
package rest

import (
	"encoding/json"
	"net/http"
)

// Ok writes a JSON-encoded response with HTTP 200 status code.
// If encoding fails, it responds with HTTP 500 Internal Server Error.
func Ok(w http.ResponseWriter, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}
