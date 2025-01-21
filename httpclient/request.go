package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
)

// RequestOption defines the function signature for request options
type RequestOption func(*http.Request) error

// WithQueryParams adds query parameters to the request
func WithQueryParams(params map[string]string) RequestOption {
	return func(req *http.Request) error {
		q := req.URL.Query()
		for k, v := range params {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
		return nil
	}
}

// WithJSONBody adds a JSON body to the request
func WithJSONBody(body interface{}) RequestOption {
	return func(req *http.Request) error {
		bodyBytes, err := json.Marshal(body)
		if err != nil {
			return err
		}
		req.Body = io.NopCloser(bytes.NewReader(bodyBytes))
		req.Header.Set("Content-Type", "application/json")
		return nil
	}
}

// WithRequestHeader adds a header to the request
func WithRequestHeader(key, value string) RequestOption {
	return func(req *http.Request) error {
		req.Header.Set(key, value)
		return nil
	}
}

func (c *Client) buildRequest(ctx context.Context, method, urlStr string, opts ...RequestOption) (*http.Request, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	if c.baseURL != "" {
		baseURL, err := url.Parse(c.baseURL)
		if err != nil {
			return nil, err
		}
		u.Scheme = baseURL.Scheme
		u.Host = baseURL.Host
		u.Path = path.Join(baseURL.Path, u.Path)
	}

	req, err := http.NewRequestWithContext(ctx, method, u.String(), nil)
	if err != nil {
		return nil, err
	}

	// Add default headers
	for k, v := range c.headers {
		req.Header.Set(k, v)
	}

	// Apply request options
	for _, opt := range opts {
		if err := opt(req); err != nil {
			return nil, err
		}
	}

	return req, nil
}
