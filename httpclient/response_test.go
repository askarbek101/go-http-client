package httpclient

import (
	"bytes"
	"io"
	"net/http"
	"testing"
)

func TestResponseMethods(t *testing.T) {
	testBody := []byte(`{"message":"test"}`)
	mockResp := &http.Response{
		Body: io.NopCloser(bytes.NewReader(testBody)),
	}

	resp := newResponse(mockResp)

	t.Run("test Bytes()", func(t *testing.T) {
		body, err := resp.Bytes()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if !bytes.Equal(body, testBody) {
			t.Errorf("Expected body %s, got %s", testBody, body)
		}
	})

	t.Run("test String()", func(t *testing.T) {
		str, err := resp.String()
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if str != string(testBody) {
			t.Errorf("Expected string %s, got %s", string(testBody), str)
		}
	})

	t.Run("test JSON()", func(t *testing.T) {
		var result struct {
			Message string `json:"message"`
		}
		err := resp.JSON(&result)
		if err != nil {
			t.Fatalf("Unexpected error: %v", err)
		}
		if result.Message != "test" {
			t.Errorf("Expected message 'test', got '%s'", result.Message)
		}
	})
}
