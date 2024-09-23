package api

import (
	api "github.com/b-url/burl/api/v1"
)

var _ api.ServerInterface = Server{}

// Server implements v1.ServerInterface.
type Server struct{}

// NewServer returns a new Server.
func NewServer() Server {
	return Server{}
}
