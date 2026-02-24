package httpclient

import (
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
)

func TestSharedClientHasTimeout(t *testing.T) {
	if SharedClient.Timeout == 0 {
		t.Fatal("SharedClient should have a non-zero Timeout")
	}
}

func TestMakeRequest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("hello"))
	}))
	defer srv.Close()

	// Override the shared client's transport to use the test server's client
	// transport so that requests go to the test server.
	original := SharedClient
	SharedClient = srv.Client()
	SharedClient.Timeout = original.Timeout
	defer func() { SharedClient = original }()

	body, err := MakeRequest(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if body != "hello" {
		t.Fatalf("expected %q, got %q", "hello", body)
	}
}

func TestMakeRequestWithClient(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("world"))
	}))
	defer srv.Close()

	body, err := MakeRequestWithClient(srv.Client(), srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if body != "world" {
		t.Fatalf("expected %q, got %q", "world", body)
	}
}

func TestMultipleRequestsReuseConnections(t *testing.T) {
	var count atomic.Int64
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		count.Add(1)
		_, _ = w.Write([]byte("ok"))
	}))
	defer srv.Close()

	client := srv.Client()

	for i := range 5 {
		body, err := MakeRequestWithClient(client, srv.URL)
		if err != nil {
			t.Fatalf("request %d: unexpected error: %v", i, err)
		}
		if body != "ok" {
			t.Fatalf("request %d: expected %q, got %q", i, "ok", body)
		}
	}

	if count.Load() != 5 {
		t.Fatalf("expected 5 requests, got %d", count.Load())
	}
}

func TestBadPatternDemoStillWorks(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("bad"))
	}))
	defer srv.Close()

	body, err := BadPatternDemo(srv.URL)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if body != "bad" {
		t.Fatalf("expected %q, got %q", "bad", body)
	}
}
