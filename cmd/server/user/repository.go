package user

import (
	"context"

	"github.com/b-url/burl/cmd/server/dbc"
	"github.com/google/uuid"
)

// Repository provides access to the user storage.
type Repository interface {
	Create(context.Context, User) error
	Get(context.Context, uuid.UUID) (User, error)
	GetByEmail(context.Context, string) (User, error)
	Update(context.Context, User) error
	Delete(context.Context, uuid.UUID) error
}

// SQLRepository represents a SQL repository.
type SQLRepository struct {
	dbc dbc.Conn
}

// NewSQLRepository creates a new SQLRepository.
func NewSQLRepository(dbc dbc.Conn) *SQLRepository {
	return &SQLRepository{dbc: dbc}
}

// Create creates a new user.
func (r *SQLRepository) Create(ctx context.Context, u User) error {
	_, err := r.dbc.ExecContext(ctx, `
		INSERT INTO users (id, username, email, create_time, update_time)
		VALUES ($1, $2, $3, $4, $5)
	`, u.ID, u.Username, u.Email, u.CreateTime, u.UpdateTime)
	return err
}

// Get returns the user with the specified ID.
func (r *SQLRepository) Get(ctx context.Context, id uuid.UUID) (User, error) {
	var u User
	err := r.dbc.QueryRowContext(ctx, `
		SELECT id, username, email, create_time, update_time
		FROM users
		WHERE id = $1
	`, id).Scan(&u.ID, &u.Username, &u.Email, &u.CreateTime, &u.UpdateTime)
	return u, err
}

// GetByEmail returns the user with the specified email.
func (r *SQLRepository) GetByEmail(ctx context.Context, email string) (User, error) {
	var u User
	err := r.dbc.QueryRowContext(ctx, `
		SELECT id, username, email, create_time, update_time
		FROM users
		WHERE email = $1
	`, email).Scan(&u.ID, &u.Username, &u.Email, &u.CreateTime, &u.UpdateTime)
	return u, err
}

// Update updates the user.
func (r *SQLRepository) Update(ctx context.Context, u User) error {
	_, err := r.dbc.ExecContext(ctx, `
		UPDATE users
		SET username = $1, email = $2, update_time = $3
		WHERE id = $4
	`, u.Username, u.Email, u.UpdateTime, u.ID)
	return err
}

// Delete deletes the user with the specified ID.
func (r *SQLRepository) Delete(ctx context.Context, id uuid.UUID) error {
	_, err := r.dbc.ExecContext(ctx, `
		DELETE FROM users
		WHERE id = $1
	`, id)
	return err
}
