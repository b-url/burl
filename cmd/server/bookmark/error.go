package bookmark

import "github.com/pkg/errors"

var (
	// ErrBookmarkRequired is returned when a bookmark is required but not provided.
	ErrBookmarkRequired = errors.New("bookmark is required")
)
