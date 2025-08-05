package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		return "", errors.New("malformed authorization header")
	}

	return splitAuth[1], nil
}

func TestGetAPIKey_Success(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/some-endpoint", nil)
    expectedKey := "secret-key-123"
    req.Header.Set("X-API-Key", expectedKey)

    key, err := GetAPIKey(req)
    if err != nil {
        t.Fatalf("expected no error, got %v", err)
    }

    if key != expectedKey {
        t.Errorf("expected key %q, got %q", expectedKey, key)
    }
}

func TestGetAPIKey_MissingHeader(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/some-endpoint", nil)

    key, err := GetAPIKey(req)
    if err == nil {
        t.Fatal("expected error, got nil")
    }

    if key != "" {
        t.Errorf("expected empty key, got %q", key)
    }

    if !errors.Is(err, ErrNoAPIKey) {
        t.Errorf("expected error %v, got %v", ErrNoAPIKey, err)
    }
}

