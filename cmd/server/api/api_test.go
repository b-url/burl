package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/b-url/burl/cmd/server/api"
	"github.com/google/go-cmp/cmp"
)

func TestNewServer(t *testing.T) {
	t.Run("NewServer returns a new Server", func(t *testing.T) {
		expected := api.Server{}
		if diff := cmp.Diff(api.NewServer(), expected); diff != "" {
			t.Errorf("NewServer() mismatch (-want +got):\n%s", diff)
		}
	})
}

func TestServer_BookmarksCreate(t *testing.T) {
	t.Run("BookmarksCreate writes a response", func(t *testing.T) {
		// Create a new Server instance
		s := api.NewServer()

		// Create a new HTTP request
		req, err := http.NewRequest(http.MethodPost, "/bookmarks", nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}

		// Create a ResponseRecorder to capture the response
		rr := httptest.NewRecorder()

		// Call the BookmarksCreate method
		s.BookmarksCreate(rr, req, "user123", "collection456")

		// Check the status code
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		// Check the response body
		expected := "created"
		if rr.Body.String() != expected {
			t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
		}
	})
}
