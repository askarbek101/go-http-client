package httpclient

import (
	"net/http"
	"time"
)

// WithBaseURL sets the base URL for the client
func WithBaseURL(baseURL string) Option {
	return func(c *Client) {
		c.baseURL = baseURL
	}
}

// WithTimeout sets the timeout for the client
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}

// WithHeader sets a default header for all requests
func WithHeader(key, value string) Option {
	return func(c *Client) {
		c.headers[key] = value
	}
}

// WithHTTPClient sets a custom http.Client
func WithHTTPClient(httpClient *http.Client) Option {
	return func(c *Client) {
		c.httpClient = httpClient
	}
}
