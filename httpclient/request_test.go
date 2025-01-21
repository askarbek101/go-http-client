package httpclient

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithQueryParams(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Get("key") != "value" {
			t.Errorf("Expected query param 'key=value', got '%s'", r.URL.Query().Get("key"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	_, err := client.Get(context.Background(), "/test",
		WithQueryParams(map[string]string{
			"key": "value",
		}),
	)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestWithJSONBody(t *testing.T) {
	type TestBody struct {
		Message string `json:"message"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type: application/json, got %s", r.Header.Get("Content-Type"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	_, err := client.Post(context.Background(), "/test",
		WithJSONBody(TestBody{Message: "test"}),
	)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestWithRequestHeader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("X-Custom") != "test" {
			t.Errorf("Expected X-Custom: test, got %s", r.Header.Get("X-Custom"))
		}
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	_, err := client.Get(context.Background(), "/test",
		WithRequestHeader("X-Custom", "test"),
	)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}
