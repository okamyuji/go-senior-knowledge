// Package httpclient demonstrates the importance of reusing http.Client
// instances in Go. Creating a new http.Client for every request causes
// connection and goroutine leaks because each client maintains its own
// connection pool and transport.
//
// The standard http.DefaultClient is shared across the entire process, but
// it has no timeout set by default. Production code should create a
// package-level client with an explicit Timeout.
package httpclient

import (
	"io"
	"net/http"
	"time"
)

// SharedClient is a package-level http.Client configured with a 10-second
// timeout. All functions in this package reuse this single client, which
// allows the underlying transport to pool and reuse TCP connections across
// multiple requests.
var SharedClient = &http.Client{
	Timeout: 10 * time.Second,
}

// MakeRequest performs a GET request to the given URL using the shared
// client. It reads and returns the response body as a string.
// The response body is always closed via defer to prevent resource leaks.
func MakeRequest(url string) (string, error) {
	resp, err := SharedClient.Get(url)
	if err != nil {
		return "", err
	}
	// ResponseBodyClose pattern: always defer resp.Body.Close() immediately
	// after checking the error from the HTTP call. Failing to close the body
	// prevents the underlying TCP connection from being returned to the pool,
	// causing connection leaks under load.
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// MakeRequestWithClient performs a GET request using the provided client.
// This is useful for testing with custom transports or httptest servers.
func MakeRequestWithClient(client *http.Client, url string) (string, error) {
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

// BadPattern demonstrates what NOT to do. Creating a new http.Client for
// every request means each call allocates a fresh http.Transport with its
// own connection pool. Connections opened by one-off clients are never
// reused, and idle connections linger until they time out, leaking file
// descriptors and goroutines.
//
//	// BAD: new client per request
//	func BadPattern(url string) (string, error) {
//	    client := &http.Client{Timeout: 10 * time.Second}
//	    resp, err := client.Get(url)
//	    ...
//	}
//
// The correct approach is to create a single http.Client at package level
// or in main() and pass it to all functions that need it.

// BadPatternDemo creates a new http.Client per call to illustrate the
// anti-pattern. In production code, this wastes resources because the
// transport's connection pool is discarded after every request.
func BadPatternDemo(url string) (string, error) {
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
