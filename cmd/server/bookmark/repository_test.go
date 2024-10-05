package bookmark_test

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/b-url/burl/cmd/server/bookmark"
	"github.com/b-url/burl/cmd/server/database"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

//nolint:tparallel // This test requires a database connection.
func TestIntegration_BookmarkRepository(t *testing.T) {
	t.Parallel()

	if testing.Short() {
		fmt.Println("skipping integration test")
		t.Skip()
		return
	}

	db := getDatabase(t)
	t.Cleanup(func() { db.Close() })

	repo := bookmark.NewRepository(db)

	// Arrange
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	require.NoError(t, err)

	// Insert new user.
	userID := uuid.MustParse("2f618d2d-e65d-4541-b09f-1cf58edc36b4")
	newUser(t, tx, userID)

	// Insert new collection
	collectionID := uuid.MustParse("c7596844-695b-400d-892c-cb1b362b8101")
	newCollection(t, tx, collectionID)

	expectedBookmark := bookmark.Bookmark{
		ID:           uuid.MustParse("ba708c9d-44d2-42aa-9585-4596d7319a51"),
		CollectionID: collectionID,
		UserID:       userID,
		URL:          "http://example.com",
		Title:        "Example",
	}

	_, err = repo.CreateBookmark(ctx, tx, expectedBookmark)
	require.NoError(t, err)

	t.Run("GetBookmark", func(t *testing.T) {
		actualBookmark, getErr := repo.GetBookmark(ctx, tx, expectedBookmark.ID, userID)
		require.NoError(t, getErr)

		assert.Equal(t, expectedBookmark.ID, actualBookmark.ID)
		assert.Equal(t, expectedBookmark.CollectionID, actualBookmark.CollectionID)
		assert.Equal(t, expectedBookmark.UserID, actualBookmark.UserID)
		assert.Equal(t, expectedBookmark.URL, actualBookmark.URL)
		assert.Equal(t, expectedBookmark.Title, actualBookmark.Title)
	})

	t.Run("non-existent bookmark", func(t *testing.T) {
		_, err = repo.GetBookmark(ctx, tx, uuid.MustParse("00000000-0000-0000-0000-000000000000"), userID)
		assert.ErrorIs(t, err, bookmark.ErrBookmarkNotFound)
	})

	t.Run("returns no result for different user", func(t *testing.T) {
		_, err = repo.GetBookmark(ctx, tx, expectedBookmark.ID, uuid.MustParse("00000000-0000-0000-0000-000000000000"))
		assert.ErrorIs(t, err, bookmark.ErrBookmarkNotFound)
	})

	err = tx.Rollback()
	require.NoError(t, err)
}

func getDatabase(t *testing.T) *sql.DB {
	config := database.Config{
		DSN: "postgres://develop:develop_secret@localhost:5432/develop?sslmode=disable",
	}

	db, err := database.NewConnection(config)
	if err != nil {
		t.Fatalf("could not connect to database: %v", err)
	}

	return db
}

func newUser(t *testing.T, tx *sql.Tx, id uuid.UUID) {
	query := `
		INSERT INTO users (id, username, email)
		VALUES ($1, $2, $3)
		ON CONFLICT DO NOTHING
	`

	_, err := tx.Exec(query, id, "testuser", "testuser@example.com")
	if err != nil {
		t.Fatalf("could not insert user: %v", err)
	}
}

func newCollection(t *testing.T, tx *sql.Tx, id uuid.UUID) {
	query := `
		INSERT INTO collections (id, user_id, name)
		VALUES ($1, $2, $3)
		ON CONFLICT DO NOTHING
	`

	_, err := tx.Exec(query, id, "2f618d2d-e65d-4541-b09f-1cf58edc36b4", "testcollection")
	if err != nil {
		t.Fatalf("could not insert collection: %v", err)
	}
}
