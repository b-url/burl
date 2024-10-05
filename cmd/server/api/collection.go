package api

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func (s *Server) CollectionsCreate(
	w http.ResponseWriter,
	_ *http.Request,
	userID uuid.UUID,
) {
	fmt.Printf("Creating a new collection for user %s\n", userID)
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *Server) CollectionsRead(
	w http.ResponseWriter,
	_ *http.Request,
	userID uuid.UUID,
	collectionID uuid.UUID,
) {
	fmt.Printf("Reading collection %s for user %s\n", collectionID, userID)
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *Server) CollectionsUpdate(
	w http.ResponseWriter,
	_ *http.Request,
	userID uuid.UUID,
	collectionID uuid.UUID,
) {
	fmt.Printf("Updating collection %s for user %s\n", collectionID, userID)
	w.WriteHeader(http.StatusNotImplemented)
}

func (s *Server) CollectionsListChildren(
	w http.ResponseWriter,
	_ *http.Request,
	userID uuid.UUID,
	collectionID uuid.UUID,
) {
	fmt.Printf("Listing children of collection %s for user %s\n", collectionID, userID)
	w.WriteHeader(http.StatusNotImplemented)
}
