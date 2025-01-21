package httpclient

import (
	"encoding/json"
	"io"
	"net/http"
)

// Response wraps the http.Response
type Response struct {
	*http.Response
	body []byte
}

func newResponse(resp *http.Response) *Response {
	return &Response{Response: resp}
}

// Bytes returns the response body as bytes
func (r *Response) Bytes() ([]byte, error) {
	if r.body != nil {
		return r.body, nil
	}

	defer r.Body.Close()
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	r.body = body
	return body, nil
}

// String returns the response body as string
func (r *Response) String() (string, error) {
	bytes, err := r.Bytes()
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

// JSON unmarshals the response body into the provided interface
func (r *Response) JSON(v interface{}) error {
	bytes, err := r.Bytes()
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, v)
}
