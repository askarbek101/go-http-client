# HTTP Client Library

A flexible and feature-rich HTTP client wrapper for Go applications that provides a clean, fluent interface for making HTTP requests.

## Features

- Simple and intuitive API for HTTP requests (GET, POST, PUT, DELETE)
- Context support for request cancellation and timeouts
- Customizable request options (headers, query parameters, body)
- Configurable client options (timeout, base URL, default headers)
- Response handling with convenient methods
- Error handling and custom error types

## Installation

```bash
go get github.com/askarbek101/go-http-client
```

## Quick Start

```go
package main

import (
    "context"
    "log"
    "time"
    "github.com/askarbek101/go-http-client"
)

func main() {
    // Create a new client with options
    client := httpclient.New(
        httpclient.WithTimeout(10 * time.Second),
        httpclient.WithBaseURL("https://api.example.com"),
    )

    // Make a GET request
    resp, err := client.Get(context.Background(), "/users", 
        httpclient.WithHeader("Authorization", "Bearer token"),
        httpclient.WithQueryParam("page", "1"),
    )
    if err != nil {
        log.Fatal(err)
    }

    // Process the response
    var users []User
    if err := resp.JSON(&users); err != nil {
        log.Fatal(err)
    }
}
```

## Client Options

The client can be configured with various options:

- `WithTimeout(duration)`: Set the client timeout
- `WithBaseURL(url)`: Set the base URL for all requests
- `WithDefaultHeaders(headers)`: Set default headers for all requests

## Request Options

Each request can be customized with:

- `WithHeader(key, value)`: Add a header to the request
- `WithHeaders(headers)`: Add multiple headers
- `WithQueryParam(key, value)`: Add a query parameter
- `WithBody(body)`: Set the request body
- `WithJSON(data)`: Set JSON request body

## Response Handling

The Response object provides methods for handling the response:

- `JSON(v interface{})`: Unmarshal response body into a struct
- `String()`: Get response body as string
- `Bytes()`: Get raw response body
- `StatusCode()`: Get HTTP status code

## Error Handling

The library provides custom error types for different scenarios:

- Request creation errors
- HTTP transport errors
- Response parsing errors
- Context cancellation

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
