package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/b-url/burl/cmd/server/bookmark"
)

type Bookmarker interface {
	CreateBookmark(ctx context.Context, b *bookmark.Bookmark) (*bookmark.Bookmark, error)
}

func (s *Server) BookmarksCreate(
	w http.ResponseWriter,
	_ *http.Request,
	userID string,
	collectionID string,
) {
	fmt.Printf("Creating a new bookmark for user %s in collection %s\n", userID, collectionID)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("created")); err != nil {
		fmt.Print(err)
	}
}

func (s *Server) BookmarksRead(
	w http.ResponseWriter,
	_ *http.Request,
	userID string,
	collectionID string,
	bookmarkID string,
) {
	fmt.Printf("Reading bookmark %s for user %s in collection %s\n", bookmarkID, userID, collectionID)
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *Server) BookmarksUpdate(
	w http.ResponseWriter,
	_ *http.Request,
	userID string,
	collectionID string,
	bookmarkID string,
) {
	fmt.Printf("Updating bookmark %s for user %s in collection %s\n", bookmarkID, userID, collectionID)
	w.WriteHeader(http.StatusNotImplemented)
}
