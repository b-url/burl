package bookmark

import (
	"context"
	"time"

	"github.com/google/uuid"
)

// Collection reprseents a collection of bookmarks.
type Collection struct {
	ID         uuid.UUID
	Name       string
	ParentID   *uuid.UUID
	UserID     uuid.UUID
	CreateTime *time.Time
	UpdateTime *time.Time
}

type CreateCollectionParams struct {
	Name     string
	ParentID *uuid.UUID
	UserID   uuid.UUID
}

func (b *Bookmarker) CreateCollection(_ context.Context, _ CreateBookmarkParams) (Collection, error) {
	return Collection{}, nil
}
