package bookmark

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

var (
	// ErrBookmarkNotFound is returned when a bookmark is not found.
	ErrBookmarkNotFound = errors.New("bookmark not found")
)

type SQLRepository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *SQLRepository {
	return &SQLRepository{DB: db}
}

// Transactionally executes a function within a database transaction. It commits the transaction
// if the function succeeds, otherwise it rolls back. If rollback fails, both errors are returned.
func (r *SQLRepository) Transactionally(ctx context.Context, f func(tx *sql.Tx) error) (err error) {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				err = fmt.Errorf("panic occurred: %v, rollback error: %w", p, rbErr)
			} else {
				err = fmt.Errorf("panic occurred: %v", p)
			}
		}
	}()

	err = f(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("transaction rollback error: %w, original error: %w", rbErr, err)
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

// CreateBookmark creates a new bookmark.
func (r *SQLRepository) CreateBookmark(ctx context.Context, tx *sql.Tx, bookmark Bookmark) (Bookmark, error) {
	query := `
		INSERT INTO bookmarks (id, collection_id, user_id, url, title)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, create_time, update_time
	`
	row := tx.QueryRowContext(
		ctx,
		query,
		bookmark.ID,
		bookmark.CollectionID,
		bookmark.UserID,
		bookmark.URL,
		bookmark.Title,
	)

	err := row.Scan(&bookmark.ID, &bookmark.CreateTime, &bookmark.UpdateTime)
	if err != nil {
		return Bookmark{}, err
	}

	return bookmark, nil
}

// GetBookmark retrieves a bookmark by its ID.
func (r *SQLRepository) GetBookmark(ctx context.Context, tx *sql.Tx, id, userID uuid.UUID) (Bookmark, error) {
	query := `
		SELECT id, collection_id, user_id, url, title, create_time, update_time
		FROM bookmarks
		WHERE id = $1 AND user_id = $2
	`
	row := tx.QueryRowContext(ctx, query, id, userID)

	bookmark := Bookmark{}
	err := row.Scan(
		&bookmark.ID,
		&bookmark.CollectionID,
		&bookmark.UserID,
		&bookmark.URL,
		&bookmark.Title,
		&bookmark.CreateTime,
		&bookmark.UpdateTime,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return Bookmark{}, ErrBookmarkNotFound
		}
		return Bookmark{}, err
	}

	return bookmark, nil
}

// CreateCollection creates a new collection in the collection table.
func (r *SQLRepository) CreateCollection(ctx context.Context, tx *sql.Tx, collection Collection) (Collection, error) {
	query := `
		INSERT INTO collections (id, user_id, name, parent_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, create_time, update_time
	`
	row := tx.QueryRowContext(
		ctx,
		query,
		collection.ID,
		collection.UserID,
		collection.Name,
		collection.ParentID,
	)

	err := row.Scan(&collection.ID, &collection.CreateTime, &collection.UpdateTime)
	if err != nil {
		return Collection{}, err
	}

	return collection, nil
}

// GetCollection retrieves a collection by its ID.
func (r *SQLRepository) GetCollection(ctx context.Context, tx *sql.Tx, id, userID uuid.UUID) (Collection, error) {
	query := `
		SELECT id, user_id, name, parent_id, create_time, update_time
		FROM collections
		WHERE id = $1 AND user_id = $2
	`
	row := tx.QueryRowContext(ctx, query, id, userID)

	collection := Collection{}
	err := row.Scan(
		&collection.ID,
		&collection.UserID,
		&collection.Name,
		&collection.ParentID,
		&collection.CreateTime,
		&collection.UpdateTime,
	)
	if err != nil {
		return Collection{}, err
	}

	return collection, nil
}
