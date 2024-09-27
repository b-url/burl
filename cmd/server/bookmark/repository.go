package bookmark

import (
	"context"
	"database/sql"
	"fmt"
)

type Repository struct {
	DB *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{DB: db}
}

// Transactionally executes a function within a database transaction. It commits the transaction
// if the function succeeds, otherwise it rolls back. If rollback fails, both errors are returned.
func (r *Repository) Transactionally(ctx context.Context, f func(tx *sql.Tx) error) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

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
func (r *Repository) CreateBookmark(ctx context.Context, tx *sql.Tx, bookmark *Bookmark) (*Bookmark, error) {
	query := `
		INSERT INTO bookmarks (collection_id, user_id, url, title)
		VALUES ($1, $2, $3, $4)
		RETURNING id, create_time, update_time
	`
	row := tx.QueryRowContext(
		ctx,
		query,
		bookmark.CollectionID,
		bookmark.UserID,
		bookmark.URL,
		bookmark.Title,
	)

	var created Bookmark
	err := row.Scan(&created.ID, &created.CreateTime, &created.UpdateTime)
	if err != nil {
		return nil, err
	}

	return &created, nil
}