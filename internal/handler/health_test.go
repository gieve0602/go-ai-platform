package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	Health(rec, req)
	if rec.Code != http.StatusOK {
		t.Fatalf("expected status code 200, but got %d\n", rec.Code)
	}
	var respHealth HealthResponse

	if err := json.Unmarshal(rec.Body.Bytes(), &respHealth); err != nil {
		t.Fatalf("invalid response: %v", respHealth)
	}

	if respHealth.Status != "ok" {
		t.Fatalf("expected json result status ok, but got %v", respHealth)
	}
}

func TestHealthMethodNotAllowed(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/health", nil)
	rec := httptest.NewRecorder()

	Health(rec, req)

	if rec.Code != http.StatusMethodNotAllowed {
		t.Fatalf("expected status code 405, but got %d\n", rec.Code)
	}
}
