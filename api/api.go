
package api

import (
	"encoding/json"
	"net/http"
)

//
type ArticleParams struct {
	Author string
	Username string
}

type ArticleResponse struct { // Corrected struct name
	Code            int   // response 2xx
	NumberOfArticles int64
}

type Error struct {
	Code    int    // response of 4xx / 5xx
	Message string
}

// writeError sends a JSON response with the specified error message and code.
func writeError(w http.ResponseWriter, message string, code int) {
	resp := Error{
		Code:    code,
		Message: message,
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(resp)
}

var (
	// Handles client-related request errors
	RequestErrorHandler = func(w http.ResponseWriter, err error) {
		writeError(w, err.Error(), http.StatusBadRequest) // Changed to StatusBadRequest
	}

	// Handles server-related internal errors
	InternalErrorHandler = func(w http.ResponseWriter) {
		writeError(w, "An unexpected error occurred.", http.StatusInternalServerError)
	}
)

