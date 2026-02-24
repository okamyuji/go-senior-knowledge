package httptransport

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestDefaultTransportConfig(t *testing.T) {
	cfg := DefaultTransportConfig()
	if cfg.MaxIdleConns != 100 {
		t.Errorf("MaxIdleConns = %d, want 100", cfg.MaxIdleConns)
	}
	if cfg.MaxIdleConnsPerHost != 10 {
		t.Errorf("MaxIdleConnsPerHost = %d, want 10", cfg.MaxIdleConnsPerHost)
	}
	if cfg.IdleConnTimeout != 90*time.Second {
		t.Errorf("IdleConnTimeout = %v, want 90s", cfg.IdleConnTimeout)
	}
}

func TestCustomTransportSettings(t *testing.T) {
	cfg := TransportConfig{
		MaxIdleConns:          50,
		MaxIdleConnsPerHost:   5,
		IdleConnTimeout:       30 * time.Second,
		DialTimeout:           3 * time.Second,
		TLSHandshakeTimeout:   3 * time.Second,
		ResponseHeaderTimeout: 5 * time.Second,
	}
	tr := CustomTransport(cfg)

	if tr.MaxIdleConns != 50 {
		t.Errorf("MaxIdleConns = %d, want 50", tr.MaxIdleConns)
	}
	if tr.MaxIdleConnsPerHost != 5 {
		t.Errorf("MaxIdleConnsPerHost = %d, want 5", tr.MaxIdleConnsPerHost)
	}
	if tr.IdleConnTimeout != 30*time.Second {
		t.Errorf("IdleConnTimeout = %v, want 30s", tr.IdleConnTimeout)
	}
	if tr.TLSHandshakeTimeout != 3*time.Second {
		t.Errorf("TLSHandshakeTimeout = %v, want 3s", tr.TLSHandshakeTimeout)
	}
	if tr.ResponseHeaderTimeout != 5*time.Second {
		t.Errorf("ResponseHeaderTimeout = %v, want 5s", tr.ResponseHeaderTimeout)
	}
}

func TestClientWithTransport(t *testing.T) {
	tr := CustomTransport(DefaultTransportConfig())
	client := ClientWithTransport(tr, 30*time.Second)
	if client.Timeout != 30*time.Second {
		t.Errorf("Client.Timeout = %v, want 30s", client.Timeout)
	}
	if client.Transport != tr {
		t.Error("Client.Transport does not match provided transport")
	}
}

func TestRequestWithCustomTransport(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Custom", "test-value")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("hello"))
	}))
	defer server.Close()

	tr := CustomTransport(DefaultTransportConfig())
	client := ClientWithTransport(tr, 10*time.Second)

	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("status = %d, want %d", resp.StatusCode, http.StatusOK)
	}
	if got := resp.Header.Get("X-Custom"); got != "test-value" {
		t.Errorf("X-Custom header = %q, want %q", got, "test-value")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	if string(body) != "hello" {
		t.Errorf("body = %q, want %q", body, "hello")
	}
}

func TestMultipleRequests(t *testing.T) {
	var requestCount int
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	tr := CustomTransport(DefaultTransportConfig())
	client := ClientWithTransport(tr, 10*time.Second)

	for i := range 5 {
		resp, err := client.Get(server.URL)
		if err != nil {
			t.Fatalf("request %d: %v", i, err)
		}
		_ = resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("request %d: status = %d, want %d", i, resp.StatusCode, http.StatusOK)
		}
	}
	if requestCount != 5 {
		t.Errorf("server received %d requests, want 5", requestCount)
	}
}

func TestTimeoutConfig(t *testing.T) {
	client := TimeoutConfig(
		2*time.Second,
		3*time.Second,
		5*time.Second,
		15*time.Second,
	)
	if client.Timeout != 15*time.Second {
		t.Errorf("Client.Timeout = %v, want 15s", client.Timeout)
	}
	tr, ok := client.Transport.(*http.Transport)
	if !ok {
		t.Fatal("Transport is not *http.Transport")
	}
	if tr.TLSHandshakeTimeout != 3*time.Second {
		t.Errorf("TLSHandshakeTimeout = %v, want 3s", tr.TLSHandshakeTimeout)
	}
	if tr.ResponseHeaderTimeout != 5*time.Second {
		t.Errorf("ResponseHeaderTimeout = %v, want 5s", tr.ResponseHeaderTimeout)
	}
}

func TestTimeoutConfigWithServer(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	}))
	defer server.Close()

	client := TimeoutConfig(
		5*time.Second,
		5*time.Second,
		5*time.Second,
		10*time.Second,
	)
	resp, err := client.Get(server.URL)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	defer func() { _ = resp.Body.Close() }()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("ReadAll: %v", err)
	}
	if string(body) != "ok" {
		t.Errorf("body = %q, want %q", body, "ok")
	}
}
