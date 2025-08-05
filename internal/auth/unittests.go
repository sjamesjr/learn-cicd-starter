
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

