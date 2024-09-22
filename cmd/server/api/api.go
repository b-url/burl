package api

import (
	"fmt"
	"net/http"

	api "github.com/b-url/burl/api/v1"
)

var _ api.ServerInterface = Server{}

// Server implements v1.ServerInterface.
type Server struct{}

// NewServer returns a new Server.
func NewServer() Server {
	return Server{}
}

// BookmarksCreate creates a new bookmark and returns it.
func (s Server) BookmarksCreate(w http.ResponseWriter, _ *http.Request, userID string, collectionID string) {
	fmt.Printf("Creating a new bookmark for user %s in collection %s\n", userID, collectionID)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write([]byte("created")); err != nil {
		fmt.Print(err)
	}
}
