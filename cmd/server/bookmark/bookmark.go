// Package bookmark contains the bookmark model and related functionality.
package bookmark

import (
	"context"
	"log/slog"
	"time"

	"github.com/b-url/burl/cmd/server/database"
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
	transactionManager *database.TransactionManager
	logger             *slog.Logger
}

func NewBookmarker(tm *database.TransactionManager, logger *slog.Logger) *Bookmarker {
	return &Bookmarker{
		transactionManager: tm,
		logger:             logger,
	}
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
	if err = b.transactionManager.Transactionally(ctx, func(tx database.Conn) error {
		repository := NewRepository(tx)

		var txErr error
		createdBookmark, txErr = repository.CreateBookmark(ctx, bookmark)
		if txErr != nil {
			return txErr
		}

		return nil
	}); err != nil {
		return createdBookmark, err
	}

	return createdBookmark, nil
}

// GetBookmark retrieves a bookmark by its ID and user ID.
func (b *Bookmarker) GetBookmark(ctx context.Context, id, userID uuid.UUID) (Bookmark, error) {
	repository := NewRepository(b.transactionManager.Database())
	// TODO: Error mapping of databse errors.
	return repository.GetBookmark(ctx, id, userID)
}
