package bookmark_test

import (
	"database/sql"
	"testing"

	"github.com/b-url/burl/cmd/server/bookmark"
)

func NewBookmarker(t *testing.T, db *sql.DB) *bookmark.Bookmarker {
	t.Helper()

	repository := bookmark.NewRepository(db)
	return bookmark.NewBookmarker(repository)
}

/*func Test_Bookmarker_CreateBookmark(t *testing.T) {
	t.Parallel()

	if !integration.RequireIntegration(t) {
		return
	}

	db, err := database.NewConnection(database.Config{
		DSN: "postgres://develop:develop_secret@database:5432/develop?sslmode=disable",
	})
	if err != nil {
		t.Fatalf("failed to create database connection: %v", err)
	}

	defer db.Close()

	// Create user.
	_, err = db.Exec("INSERT INTO users (username, email) VALUES ('test', 'test@test.com')")
	if err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	_ = NewBookmarker(t, db)
}*/
