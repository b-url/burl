// Package bookmark contains the bookmark model and related functionality.
package bookmark

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

// Bookmark represents a saved URL with additional metadata.
// A bookmark can optionally be associated with a collection and a user.
// If the bookmark does not belong to a collection, it implies that it is a top-level bookmark.
type Bookmark struct {
	ID           uuid.UUID
	CollectionID uuid.UUID
	UserID       uuid.UUID
	URL          string
	Title        string
	CreateTime   *time.Time
	UpdateTime   *time.Time
}

// Bookmarker is responsible for bookmark-related operations.
// It encapsulates the bookmark repository and perform related side effects.
type Bookmarker struct {
	repository Repository
}

type Repository interface {
	Transactionally(ctx context.Context, f func(tx *sql.Tx) error) (err error)
	CreateBookmark(ctx context.Context, tx *sql.Tx, bookmark Bookmark) (Bookmark, error)
	GetBookmark(ctx context.Context, tx *sql.Tx, id, userID uuid.UUID) (Bookmark, error)
}

func NewBookmarker(repository Repository) *Bookmarker {
	return &Bookmarker{repository: repository}
}

type CreateBookmarkParams struct {
	URL          string
	Title        string
	CollectionID uuid.UUID
	UserID       uuid.UUID
	Tags         []string
}

// TODO: Handle tag upsert.
// CreateBookmark creates a new bookmark.
func (b *Bookmarker) CreateBookmark(ctx context.Context, params CreateBookmarkParams) (Bookmark, error) {
	createdBookmark := Bookmark{}

	id, err := uuid.NewV7()
	if err != nil {
		return createdBookmark, err
	}

	bookmark := Bookmark{
		ID:           id,
		CollectionID: params.CollectionID,
		URL:          params.URL,
		Title:        params.Title,
		UserID:       params.UserID,
	}
	if err = b.repository.Transactionally(ctx, func(tx *sql.Tx) error {
		createdBookmark, err = b.repository.CreateBookmark(ctx, tx, bookmark)
		if err != nil {
			return err
		}

		return nil
	}); err != nil {
		return createdBookmark, err
	}

	return createdBookmark, nil
}

// GetBookmark retrieves a bookmark by its ID and user ID.
func (b *Bookmarker) GetBookmark(ctx context.Context, id, userID uuid.UUID) (Bookmark, error) {
	bookmark := Bookmark{}
	err := b.repository.Transactionally(ctx, func(tx *sql.Tx) error {
		var err error
		bookmark, err = b.repository.GetBookmark(ctx, tx, id, userID)
		if err != nil {
			// TODO: Marshal error model.
			return err
		}

		return nil
	})
	if err != nil {
		return bookmark, err
	}

	return bookmark, nil
}
