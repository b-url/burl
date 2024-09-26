// Package bookmark contains the bookmark model and related functionality.
package bookmark

import (
	"context"
	"database/sql"
	"time"
)

// Bookmark represents a saved URL with additional metadata.
// A bookmark can optionally be associated with a collection and a user.
// If the bookmark does not belong to a collection, it implies that it is a top-level bookmark.
type Bookmark struct {
	ID           *int64
	CollectionID int64
	UserID       int64
	URL          string
	Title        string
	Description  string
	CreateTime   time.Time
	UpdateTime   time.Time
}

// Bookmarker is responsible for bookmark-related operations.
// It encapsulates the bookmark repository and perform related side effects.
type Bookmarker struct {
	repository *Repository
}

func NewBookmarker(repository *Repository) *Bookmarker {
	return &Bookmarker{repository: repository}
}

// CreateBookmark creates a new bookmark.
func (b *Bookmarker) CreateBookmark(ctx context.Context, bookmark *Bookmark) (*Bookmark, error) {
	if bookmark == nil {
		return nil, ErrBookmarkRequired
	}

	var createdBookmark *Bookmark
	if err := b.repository.Transactionally(ctx, func(tx *sql.Tx) error {
		var err error
		createdBookmark, err = b.repository.CreateBookmark(ctx, tx, bookmark)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return nil, err
	}

	return createdBookmark, nil
}
