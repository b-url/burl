package api

import (
	"fmt"
	"net/http"

	v1 "github.com/b-url/burl/api/v1"
)

var _ v1.ServerInterface = Server{}

// Server implements v1.ServerInterface.
type Server struct{}

// NewServer returns a new Server.
func NewServer() Server {
	return Server{}
}

// BookmarksCreate creates a new bookmark and returns it.
func (s Server) BookmarksCreate(w http.ResponseWriter, r *http.Request, userId string, collectionId string) {
	fmt.Print("BookmarksCreate")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("created"))
}
