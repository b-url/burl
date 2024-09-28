package api

import (
	api "github.com/b-url/burl/api/v1"
)

var _ api.ServerInterface = &Server{}

type Server struct {
	Bookmarker Bookmarker
}

// NewServer returns a new Server.
func NewServer(b Bookmarker) *Server {
	return &Server{
		Bookmarker: b,
	}
}
