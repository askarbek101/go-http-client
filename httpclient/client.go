package httpclient

import (
	"context"
	"net/http"
	"time"
)

// Client represents the HTTP client wrapper
type Client struct {
	httpClient *http.Client
	baseURL    string
	headers    map[string]string
	timeout    time.Duration
}

// Option defines the function signature for client options
type Option func(*Client)

// New creates a new HTTP client instance with the given options
func New(options ...Option) *Client {
	client := &Client{
		httpClient: &http.Client{},
		headers:    make(map[string]string),
		timeout:    30 * time.Second,
	}

	for _, opt := range options {
		opt(client)
	}

	client.httpClient.Timeout = client.timeout
	return client
}

// Get performs a GET request
func (c *Client) Get(ctx context.Context, url string, opts ...RequestOption) (*Response, error) {
	return c.Do(ctx, http.MethodGet, url, opts...)
}

// Post performs a POST request
func (c *Client) Post(ctx context.Context, url string, opts ...RequestOption) (*Response, error) {
	return c.Do(ctx, http.MethodPost, url, opts...)
}

// Put performs a PUT request
func (c *Client) Put(ctx context.Context, url string, opts ...RequestOption) (*Response, error) {
	return c.Do(ctx, http.MethodPut, url, opts...)
}

// Delete performs a DELETE request
func (c *Client) Delete(ctx context.Context, url string, opts ...RequestOption) (*Response, error) {
	return c.Do(ctx, http.MethodDelete, url, opts...)
}

// Do performs the HTTP request
func (c *Client) Do(ctx context.Context, method, url string, opts ...RequestOption) (*Response, error) {
	req, err := c.buildRequest(ctx, method, url, opts...)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	return newResponse(resp), nil
}
