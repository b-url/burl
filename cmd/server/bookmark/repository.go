package bookmark

import (
	"context"
	"database/sql"
	"errors"

	"github.com/b-url/burl/cmd/server/database"
	"github.com/google/uuid"
)

var (
	// ErrBookmarkNotFound is returned when a bookmark is not found.
	ErrBookmarkNotFound = errors.New("bookmark not found")
)

type SQLRepository struct {
	conn database.Conn
}

func NewRepository(c database.Conn) *SQLRepository {
	return &SQLRepository{conn: c}
}

// CreateBookmark creates a new bookmark.
func (r *SQLRepository) CreateBookmark(ctx context.Context, bookmark Bookmark) (Bookmark, error) {
	query := `
		INSERT INTO bookmarks (id, collection_id, user_id, url, title)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, create_time, update_time
	`
	row := r.conn.QueryRowContext(
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
func (r *SQLRepository) GetBookmark(ctx context.Context, id, userID uuid.UUID) (Bookmark, error) {
	query := `
		SELECT id, collection_id, user_id, url, title, create_time, update_time
		FROM bookmarks
		WHERE id = $1 AND user_id = $2
	`
	row := r.conn.QueryRowContext(ctx, query, id, userID)

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
func (r *SQLRepository) CreateCollection(ctx context.Context, collection Collection) (Collection, error) {
	query := `
		INSERT INTO collections (id, user_id, name, parent_id)
		VALUES ($1, $2, $3, $4)
		RETURNING id, create_time, update_time
	`
	row := r.conn.QueryRowContext(
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
func (r *SQLRepository) GetCollection(ctx context.Context, id, userID uuid.UUID) (Collection, error) {
	query := `
		SELECT id, user_id, name, parent_id, create_time, update_time
		FROM collections
		WHERE id = $1 AND user_id = $2
	`
	row := r.conn.QueryRowContext(ctx, query, id, userID)

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

// UpdateCollection updates a collection in the collection table.
func (r *SQLRepository) UpdateCollection(ctx context.Context, collection Collection) (Collection, error) {
	query := `
		UPDATE collections
		SET name = $1, parent_id = $2
		WHERE id = $3 AND user_id = $4
	`

	_, err := r.conn.ExecContext(ctx, query, collection.Name, collection.ParentID, collection.ID, collection.UserID)
	if err != nil {
		return Collection{}, err
	}

	return collection, nil
}
