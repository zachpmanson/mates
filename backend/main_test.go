package backend
package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHealthHandler(t *testing.T) {
	startTime = time.Now()
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthHandler)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Fatalf("failed request: %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("status not OK: %d", resp.StatusCode)
	}
	var h HealthResponse
	if err := json.NewDecoder(resp.Body).Decode(&h); err != nil {
		t.Fatalf("decode: %v", err)
	}
	if h.Status != "ok" {
		t.Fatalf("expected status ok, got %s", h.Status)
	}
}
