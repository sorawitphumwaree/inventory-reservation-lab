package httpapi

import (
	"encoding/json/v2"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthHandler(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "/health", nil)
	recorder := httptest.NewRecorder()

	HealthHandler(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Errorf("status code: got %d, want %d",
			recorder.Code, http.StatusOK)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("content type: got %q, want %q",
			contentType, "application/json")
	}

	var body map[string]string

	err := json.Unmarshal(recorder.Body.Bytes(), &body)
	if err != nil {
		t.Fatalf("decode response JSON: %v", err)
	}

	if body["status"] != "ok" {
		t.Errorf("status: got %q, want %q", body["status"], "ok")
	}
}
