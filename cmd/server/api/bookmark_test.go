package api_test

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	v1 "github.com/b-url/burl/api/v1"
	"github.com/b-url/burl/cmd/server/api"
	"github.com/b-url/burl/cmd/server/bookmark"
	"github.com/google/uuid"
)

//go:generate moq -out bookmark_mock_test.go -pkg api_test -stub ../bookmark Repository

func TestServer_BookmarksCreate(t *testing.T) {
	t.Parallel()

	t.Run("BookmarksCreate writes a response", func(t *testing.T) {
		t.Parallel()

		// Arrange.
		repository := &RepositoryMock{
			CreateBookmarkFunc: func(_ context.Context, _ *sql.Tx, b bookmark.Bookmark) (bookmark.Bookmark, error) {
				return b, nil
			},
			TransactionallyFunc: func(_ context.Context, f func(tx *sql.Tx) error) error {
				return f(nil)
			},
		}
		bookmarker := bookmark.NewBookmarker(repository)
		s := api.NewServer(bookmarker)

		bookmarkCreate := v1.BookmarkCreate{
			Url:   "https://example.com",
			Title: "Example",
			Tags:  []string{"example"},
		}
		body, err := json.Marshal(bookmarkCreate)
		if err != nil {
			t.Fatalf("could not marshal bookmarkCreate: %v", err)
		}
		req, err := http.NewRequest(http.MethodPost, "/bookmarks", bytes.NewBuffer(body))
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		req.Header.Set("Content-Type", "application/json")
		rr := httptest.NewRecorder()

		// Act.
		s.BookmarksCreate(rr, req,
			uuid.MustParse("00000000-0000-0000-0000-000000000000"), uuid.MustParse("00000000-0000-0000-0000-000000000000"))

		// Assert.
		if status := rr.Code; status != http.StatusCreated {
			t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
		}

		var b v1.Bookmark
		if err = json.Unmarshal(rr.Body.Bytes(), &b); err != nil {
			t.Errorf("handler returned invalid body: %v", err)
		}
	})
}
