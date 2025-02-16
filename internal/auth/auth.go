package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetApiKey extracts an API key from the headers of an HTTP request
//Example:
//Authorization: API_KEY_dgybrtybrt...
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("No 'Authorization' header found")
	}

	vals := strings.Split(val, "API_KEY_")

	if len(vals) != 2 {
		return "", errors.New("'Authorization' found malformed, e.g Authorization: API_KEY_dgybrt...")
	}

	return strings.TrimSpace(vals[1]), nil
}
