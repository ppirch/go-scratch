package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts the API Key
// from the header of the HTTP request
// Example:
// Authorization: Bearer <API_KEY>
func GetAPIKey(r *http.Request) (string, error) {
	val := r.Header.Get("Authorization")
	if val == "" {
		return "", errors.New("no Authorization header found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid Authorization header")
	}

	if vals[0] != "Bearer" {
		return "", errors.New("invalid Authorization header")
	}

	return vals[1], nil
}
