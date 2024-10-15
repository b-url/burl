package dbc

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
)

var (
	_ Conn = (*sql.DB)(nil)
	_ Conn = (*sql.Tx)(nil)
)

// Conn represents a connection to a database.
// It is implemented by *sql.DB and *sql.Tx and can be used to perform queries and execute statements.
type Conn interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

// Manager manages the database connection.
type Manager struct {
	logger *slog.Logger
	db     *sql.DB
}

// NewManager creates a new Manager.
func NewManager(lgger *slog.Logger, db *sql.DB) *Manager {
	return &Manager{db: db}
}

// DB returns the database connection.
func (m *Manager) DB() *sql.DB {
	return m.db
}

// TX executes the function f in a transaction.
func (m *Manager) TX(ctx context.Context, f func(tx Conn) error) error {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		}
	}()

	if err := f(tx); err != nil {
		m.logger.Error("transaction error", slog.String("error", err.Error()))
		if rErr := tx.Rollback(); rErr != nil {
			return fmt.Errorf("rollback error: %w", rErr)
		}
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
