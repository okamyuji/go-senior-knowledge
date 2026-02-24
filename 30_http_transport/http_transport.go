// Package httptransport demonstrates how to configure http.Transport and
// http.Client for production use, including connection pooling, timeouts,
// and keep-alive behavior.
//
// Connection pooling: http.Transport maintains a pool of idle connections
// keyed by scheme+host. Reusing connections avoids the cost of TCP and TLS
// handshakes. MaxIdleConns limits the total pool size, MaxIdleConnsPerHost
// limits per-host idle connections (default is 2, often too low for high
// throughput to a single host).
//
// Keep-alive: HTTP/1.1 connections are kept alive by default. The transport
// closes idle connections after IdleConnTimeout. Disabling keep-alive forces
// a new connection for every request, which is useful for testing but harmful
// for performance.
//
// Timeout layers (from outermost to innermost):
//
//	http.Client.Timeout          - overall request deadline (dial + TLS + headers + body)
//	Transport.DialContext         - TCP connection establishment
//	Transport.TLSHandshakeTimeout - TLS negotiation
//	Transport.ResponseHeaderTimeout - time to receive response headers after request is sent
package httptransport

import (
	"net"
	"net/http"
	"time"
)

// TransportConfig holds parameters for constructing a customized http.Transport.
type TransportConfig struct {
	MaxIdleConns        int
	MaxIdleConnsPerHost int
	IdleConnTimeout     time.Duration
	DialTimeout         time.Duration
	TLSHandshakeTimeout time.Duration
	ResponseHeaderTimeout time.Duration
}

// DefaultTransportConfig returns a TransportConfig with sensible defaults for
// production services that communicate with a small number of backends.
func DefaultTransportConfig() TransportConfig {
	return TransportConfig{
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   10,
		IdleConnTimeout:       90 * time.Second,
		DialTimeout:           5 * time.Second,
		TLSHandshakeTimeout:   5 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
	}
}

// CustomTransport creates an http.Transport configured according to cfg.
func CustomTransport(cfg TransportConfig) *http.Transport {
	return &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: cfg.DialTimeout,
		}).DialContext,
		MaxIdleConns:          cfg.MaxIdleConns,
		MaxIdleConnsPerHost:   cfg.MaxIdleConnsPerHost,
		IdleConnTimeout:       cfg.IdleConnTimeout,
		TLSHandshakeTimeout:   cfg.TLSHandshakeTimeout,
		ResponseHeaderTimeout: cfg.ResponseHeaderTimeout,
	}
}

// ClientWithTransport creates an http.Client that uses the provided transport
// and applies an overall request timeout. The client timeout is the outermost
// deadline covering the entire request lifecycle.
func ClientWithTransport(transport *http.Transport, timeout time.Duration) *http.Client {
	return &http.Client{
		Transport: transport,
		Timeout:   timeout,
	}
}

// TimeoutConfig demonstrates layered timeout configuration by returning a
// fully configured http.Client. Each timeout layer protects against a
// different class of failure:
//   - DialTimeout: the remote host is unreachable
//   - TLSHandshakeTimeout: the TLS negotiation is slow or hung
//   - ResponseHeaderTimeout: the server accepted the connection but is slow to respond
//   - clientTimeout: the entire operation (including reading the body) is taking too long
func TimeoutConfig(dialTimeout, tlsTimeout, headerTimeout, clientTimeout time.Duration) *http.Client {
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: dialTimeout,
		}).DialContext,
		TLSHandshakeTimeout:   tlsTimeout,
		ResponseHeaderTimeout: headerTimeout,
	}
	return &http.Client{
		Transport: transport,
		Timeout:   clientTimeout,
	}
}
