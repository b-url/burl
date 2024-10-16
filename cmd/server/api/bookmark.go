package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	v1 "github.com/b-url/burl/api/v1"
	"github.com/b-url/burl/cmd/server/bookmark"
	"github.com/google/uuid"
)

type Bookmarker interface {
	CreateBookmark(ctx context.Context, b bookmark.CreateBookmarkParams) (bookmark.Bookmark, error)
	GetBookmark(ctx context.Context, id, userID uuid.UUID) (bookmark.Bookmark, error)
}

// TODO: http.Error should be replaced by the error model.
// TODO: Replace all fmt.Println with slog logging calls.

func (s *Server) BookmarksCreate(
	w http.ResponseWriter,
	r *http.Request,
	userID uuid.UUID,
	collectionID uuid.UUID,
) {
	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "failed to read request body", http.StatusBadRequest)
		return
	}

	var request v1.BookmarkCreate
	if err = json.Unmarshal(body, &request); err != nil {
		fmt.Println(err)
		http.Error(w, "failed to unmarshal request body", http.StatusBadRequest)
		return
	}

	params := bookmark.CreateBookmarkParams{
		URL:          request.Url,
		Title:        request.Title,
		Tags:         request.Tags,
		UserID:       userID,
		CollectionID: collectionID,
	}

	createdBookmark, err := s.Bookmarker.CreateBookmark(ctx, params)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	b := newBookmark(createdBookmark, collectionID)
	writeJSONResponse(w, b, http.StatusCreated)
}

func newBookmark(bookmark bookmark.Bookmark, collectionID uuid.UUID) v1.Bookmark {
	b := v1.Bookmark{
		Id:         bookmark.ID,
		ParentId:   &collectionID,
		Url:        bookmark.URL,
		Title:      bookmark.Title,
		CreateTime: bookmark.CreateTime,
		UpdateTime: bookmark.UpdateTime,
	}
	return b
}

func (s *Server) BookmarksRead(
	w http.ResponseWriter,
	r *http.Request,
	userID uuid.UUID,
	_ uuid.UUID,
	bookmarkID uuid.UUID,
) {
	ctx := r.Context()
	s.logger.DebugContext(ctx, "Reading bookmark", "bookmarkID", bookmarkID, "userID", userID)

	b, err := s.Bookmarker.GetBookmark(ctx, bookmarkID, userID)
	if err != nil {
		// Check if the error is a not found error.
		if errors.Is(err, bookmark.ErrBookmarkNotFound) {
			s.logger.WarnContext(ctx, "Bookmark not found", "bookmarkID", bookmarkID, "userID", userID)
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	apiBookmark := newBookmark(b, b.CollectionID)
	writeJSONResponse(w, apiBookmark, http.StatusOK)
}

func (s *Server) BookmarksUpdate(
	w http.ResponseWriter,
	_ *http.Request,
	userID uuid.UUID,
	collectionID uuid.UUID,
	bookmarkID uuid.UUID,
) {
	fmt.Printf("Updating bookmark %s for user %s in collection %s\n", bookmarkID, userID, collectionID)
	w.WriteHeader(http.StatusNotImplemented)
}
