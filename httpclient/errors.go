package httpclient

import "fmt"

// HTTPError represents an HTTP error response
type HTTPError struct {
	StatusCode int
	Status     string
	Body       []byte
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTP error: %d %s", e.StatusCode, e.Status)
}
