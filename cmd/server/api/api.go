package api

import (
	"encoding/json"
	"log/slog"
	"net/http"

	api "github.com/b-url/burl/api/v1"
)

var _ api.ServerInterface = &Server{}

type Server struct {
	Bookmarker Bookmarker

	logger *slog.Logger
}

// NewServer returns a new Server.
func NewServer(b Bookmarker, logger *slog.Logger) *Server {
	return &Server{
		Bookmarker: b,
		logger:     logger,
	}
}

func writeJSONResponse[T any](w http.ResponseWriter, b T, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if err := json.NewEncoder(w).Encode(b); err != nil {
		http.Error(w, "failed to encode to json", http.StatusInternalServerError)
	}
}
