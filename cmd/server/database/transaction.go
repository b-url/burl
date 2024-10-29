package database

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

// TransactionManager manages the database connection and can
// execute functions within a transaction.
type TransactionManager struct {
	logger *slog.Logger
	db     *sql.DB
}

// NewContext creates a new Manager.
func NewContext(logger *slog.Logger, db *sql.DB) *TransactionManager {
	return &TransactionManager{db: db, logger: logger}
}

func (c *TransactionManager) Database() *sql.DB {
	return c.db
}

// Transactionally executes a function within a database transaction. It commits the transaction
// if the function succeeds, otherwise it rolls back. If rollback fails, both errors are returned.
func (c *TransactionManager) Transactionally(ctx context.Context, f func(tx *sql.Tx) error) (err error) {
	tx, err := c.db.BeginTx(ctx, nil)
	if err != nil {
		c.logger.ErrorContext(ctx, "failed to begin transaction", "error", err)
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			if rbErr := tx.Rollback(); rbErr != nil {
				c.logger.ErrorContext(ctx, "panic occurred and rollback failed", "panic", p, "rollback error", rbErr)
				err = fmt.Errorf("panic occurred: %v, rollback error: %w", p, rbErr)
			} else {
				c.logger.ErrorContext(ctx, "panic occurred", "panic", p)
				err = fmt.Errorf("panic occurred: %v", p)
			}
		}
	}()

	err = f(tx)
	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			c.logger.ErrorContext(ctx, "transaction rollback error", "rollback error", rbErr, "original error", err)
			return fmt.Errorf("transaction rollback error: %w, original error: %w", rbErr, err)
		}
		c.logger.ErrorContext(ctx, "transaction function error", "error", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		c.logger.ErrorContext(ctx, "failed to commit transaction", "error", err)
		return err
	}

	return nil
}
