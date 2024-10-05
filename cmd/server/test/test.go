// Package test contains test utilities and helpers for unit tests.
package test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

// NewRequest creates a new HTTP request for testing with optional JSON body.
// Returns an HTTP response recorder and the request.
func NewRequest[T any](t *testing.T, method, url string, body *T) (*httptest.ResponseRecorder, *http.Request) {
	var buf *bytes.Buffer
	if body != nil {
		jsonBody, err := json.Marshal(body)
		if err != nil {
			t.Fatalf("could not marshal request body: %v", err)
		}
		buf = bytes.NewBuffer(jsonBody)
	} else {
		buf = bytes.NewBuffer(nil)
	}

	req, err := http.NewRequestWithContext(context.TODO(), method, url, buf)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	return rr, req
}
