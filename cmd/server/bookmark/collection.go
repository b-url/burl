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

func GetCollection(_ context.Context, _ uuid.UUID, _ uuid.UUID) (Collection, error) {
	return Collection{}, nil
}

type CreateCollectionParams struct {
	Name     string
	ParentID *uuid.UUID
	UserID   uuid.UUID
}

func (b *Bookmarker) CreateCollection(_ context.Context, _ CreateBookmarkParams) (Collection, error) {
	return Collection{}, nil
}

type UpdateCollectionParams struct {
	ID       uuid.UUID
	Name     string
	ParentID *uuid.UUID
	UserID   uuid.UUID
}

func (b *Bookmarker) UpdateCollection(_ context.Context, _ UpdateCollectionParams) (Collection, error) {
	return Collection{}, nil
}
