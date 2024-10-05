package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	v1 "github.com/b-url/burl/api/v1"
	"github.com/b-url/burl/cmd/server/bookmark"
	"github.com/google/uuid"
)

type Bookmarker interface {
	CreateBookmark(ctx context.Context, b bookmark.CreateBookmarkParams) (bookmark.Bookmark, error)
}

// TODO: http.Error should be replaced by the error model.
// TODO: Replace all fmt.Println with slog logging calls.
// TODO: Extract the conversion from bookmark.Bookmark to v1.Bookmark to a function.
// TODO: Extract the json response writing.

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

	b := v1.Bookmark{
		Id:         createdBookmark.ID,
		ParentId:   &collectionID,
		Url:        createdBookmark.URL,
		Title:      createdBookmark.Title,
		CreateTime: createdBookmark.CreateTime,
		UpdateTime: createdBookmark.UpdateTime,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	// Write response body as json
	if err = json.NewEncoder(w).Encode(b); err != nil {
		http.Error(w, "failed to encode to json", http.StatusInternalServerError)
	}
}

func (s *Server) BookmarksRead(
	w http.ResponseWriter,
	_ *http.Request,
	userID uuid.UUID,
	collectionID uuid.UUID,
	bookmarkID uuid.UUID,
) {
	fmt.Printf("Reading bookmark %s for user %s in collection %s\n", bookmarkID, userID, collectionID)
	w.WriteHeader(http.StatusNotImplemented)
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
