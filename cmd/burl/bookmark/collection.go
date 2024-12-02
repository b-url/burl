package bookmark

import (
	"context"
	"time"

	"github.com/b-url/burl/cmd/burl/database"
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

func (b *Bookmarker) CreateCollection(ctx context.Context, params CreateCollectionParams) (Collection, error) {
	creartedCollection := Collection{}

	id, err := uuid.NewV7()
	if err != nil {
		return creartedCollection, err
	}

	c := Collection{
		ID:       id,
		Name:     params.Name,
		ParentID: params.ParentID,
		UserID:   params.UserID,
	}

	if err = b.transactionManager.Transactionally(ctx, func(tx database.Conn) error {
		repository := NewRepository(tx)

		var txErr error
		creartedCollection, txErr = repository.CreateCollection(ctx, c)
		if txErr != nil {
			return txErr
		}

		return nil
	}); err != nil {
		return creartedCollection, err
	}

	return creartedCollection, nil
}

type UpdateCollectionParams struct {
	ID       uuid.UUID
	Name     string
	ParentID *uuid.UUID
	UserID   uuid.UUID
}

func (b *Bookmarker) UpdateCollection(ctx context.Context, params UpdateCollectionParams) (Collection, error) {
	updatedCollection := Collection{}

	if err := b.transactionManager.Transactionally(ctx, func(tx database.Conn) error {
		repository := NewRepository(tx)

		// Get the collection to update.
		collection, err := repository.GetCollection(ctx, params.ID, params.UserID)
		if err != nil {
			return err
		}

		// Update the collection.
		uc := Collection{
			ID:         collection.ID,
			Name:       params.Name,
			ParentID:   params.ParentID,
			UserID:     params.UserID,
			CreateTime: collection.CreateTime,
			UpdateTime: func(t time.Time) *time.Time { return &t }(time.Now()),
		}

		var txErr error
		updatedCollection, txErr = repository.UpdateCollection(ctx, uc)
		if txErr != nil {
			return txErr
		}

		return nil
	}); err != nil {
		return updatedCollection, err
	}

	return updatedCollection, nil
}
