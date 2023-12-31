package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts API Key from
// the headers of the http eequest
// Example:
// Authorization: ApiKey {insert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("no authentication info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed Auth Header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed first part of Auth Header")
	}
	return vals[1], nil
}
