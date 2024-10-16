package api_test

import (
	"context"
	"database/sql"
	"encoding/json"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	v1 "github.com/b-url/burl/api/v1"
	"github.com/b-url/burl/cmd/server/api"
	"github.com/b-url/burl/cmd/server/bookmark"
	"github.com/b-url/burl/cmd/server/test"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

//go:generate moq -out bookmark_mock_test.go -pkg api_test -stub ../bookmark Repository

func setupServerWithMock(repository *RepositoryMock) *api.Server {
	bookmarker := bookmark.NewBookmarker(repository, slog.New(slog.NewTextHandler(os.Stdout, nil)))
	return api.NewServer(bookmarker, slog.New(slog.NewTextHandler(os.Stdout, nil)))
}

func validateResponse(t *testing.T, rr *httptest.ResponseRecorder, expectedStatus int, responseBody interface{}) {
	require.Equal(t, expectedStatus, rr.Code, "handler returned wrong status code")

	err := json.Unmarshal(rr.Body.Bytes(), responseBody)
	require.NoError(t, err, "handler returned invalid body")
}

func TestServer_BookmarksCreate(t *testing.T) {
	t.Parallel()

	t.Run("BookmarksCreate writes a response", func(t *testing.T) {
		t.Parallel()

		// Arrange
		repository := &RepositoryMock{
			CreateBookmarkFunc: func(_ context.Context, _ *sql.Tx, b bookmark.Bookmark) (bookmark.Bookmark, error) {
				return b, nil
			},
			TransactionallyFunc: func(_ context.Context, f func(tx *sql.Tx) error) error {
				return f(nil)
			},
		}
		server := setupServerWithMock(repository)

		bookmarkCreate := v1.BookmarkCreate{
			Url:   "https://example.com",
			Title: "Example",
			Tags:  []string{"example"},
		}

		// Act
		rr, req := test.NewRequest(t, http.MethodPost, "/bookmarks", &bookmarkCreate)
		server.BookmarksCreate(rr,
			req, uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			uuid.MustParse("00000000-0000-0000-0000-000000000000"))

		// Assert
		var createdBookmark v1.Bookmark
		validateResponse(t, rr, http.StatusCreated, &createdBookmark)
	})
}

func TestServer_BookmarksRead(t *testing.T) {
	t.Parallel()

	t.Run("BookmarksRead writes a response", func(t *testing.T) {
		t.Parallel()

		// Arrange
		repository := &RepositoryMock{
			GetBookmarkFunc: func(_ context.Context, _ *sql.Tx, id, userID uuid.UUID) (bookmark.Bookmark, error) {
				return bookmark.Bookmark{
					ID:           id,
					CollectionID: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
					URL:          "https://example.com",
					Title:        "Example",
					UserID:       userID,
				}, nil
			},
			TransactionallyFunc: func(_ context.Context, f func(tx *sql.Tx) error) error {
				return f(nil)
			},
		}
		server := setupServerWithMock(repository)

		// Act
		rr, req := test.NewRequest[any](
			t,
			http.MethodGet,
			"/bookmarks/00000000-0000-0000-0000-000000000000",
			nil)
		server.BookmarksRead(
			rr,
			req,
			uuid.MustParse("00000000-0000-0000-0000-000000000000"),
			uuid.UUID{},
			uuid.MustParse("00000000-0000-0000-0000-000000000000"))

		// Assert
		var bookmark v1.Bookmark
		validateResponse(t, rr, http.StatusOK, &bookmark)
	})
}
