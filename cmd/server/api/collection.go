package api

import (
	"fmt"
	"net/http"
)

func (s Server) CollectionsCreate(
	w http.ResponseWriter,
	_ *http.Request,
	userID string,
) {
	fmt.Printf("Creating a new collection for user %s\n", userID)
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) CollectionsRead(
	w http.ResponseWriter,
	_ *http.Request,
	userID string,
	collectionID string,
) {
	fmt.Printf("Reading collection %s for user %s\n", collectionID, userID)
	w.WriteHeader(http.StatusNotImplemented)
}

func (s Server) CollectionsUpdate(
	w http.ResponseWriter,
	_ *http.Request,
	userID string,
	collectionID string,
) {
	fmt.Printf("Updating collection %s for user %s\n", collectionID, userID)
	w.WriteHeader(http.StatusNotImplemented)
}
