package httpclient

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClientCreation(t *testing.T) {
	client := New(
		WithBaseURL("https://api.example.com"),
		WithTimeout(5*time.Second),
		WithHeader("User-Agent", "Test/1.0"),
	)

	if client.baseURL != "https://api.example.com" {
		t.Errorf("Expected baseURL to be %s, got %s", "https://api.example.com", client.baseURL)
	}

	if client.timeout != 5*time.Second {
		t.Errorf("Expected timeout to be %v, got %v", 5*time.Second, client.timeout)
	}

	if client.headers["User-Agent"] != "Test/1.0" {
		t.Errorf("Expected User-Agent header to be %s, got %s", "Test/1.0", client.headers["User-Agent"])
	}
}

func TestHTTPMethods(t *testing.T) {
	tests := []struct {
		name           string
		method         string
		expectedMethod string
		handler        func(w http.ResponseWriter, r *http.Request)
	}{
		{
			name:           "GET request",
			method:         "GET",
			expectedMethod: "GET",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "GET" {
					t.Errorf("Expected GET request, got %s", r.Method)
				}
				w.WriteHeader(http.StatusOK)
			},
		},
		{
			name:           "POST request",
			method:         "POST",
			expectedMethod: "POST",
			handler: func(w http.ResponseWriter, r *http.Request) {
				if r.Method != "POST" {
					t.Errorf("Expected POST request, got %s", r.Method)
				}
				w.WriteHeader(http.StatusCreated)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(tt.handler))
			defer server.Close()

			client := New(WithBaseURL(server.URL))

			var resp *Response
			var err error

			ctx := context.Background()
			switch tt.method {
			case "GET":
				resp, err = client.Get(ctx, "/test")
			case "POST":
				resp, err = client.Post(ctx, "/test")
			}

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
				t.Errorf("Expected status code 200 or 201, got %d", resp.StatusCode)
			}
		})
	}
}

func TestIntegration(t *testing.T) {
	type TestResponse struct {
		Message string `json:"message"`
	}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/test/json":
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(TestResponse{Message: "success"})
		case "/test/error":
			w.WriteHeader(http.StatusBadRequest)
		}
	}))
	defer server.Close()

	client := New(WithBaseURL(server.URL))

	t.Run("successful JSON response", func(t *testing.T) {
		resp, err := client.Get(context.Background(), "/test/json")
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}

		var result TestResponse
		if err := resp.JSON(&result); err != nil {
			t.Fatalf("Failed to parse JSON: %v", err)
		}

		if result.Message != "success" {
			t.Errorf("Expected message 'success', got '%s'", result.Message)
		}
	})
}
