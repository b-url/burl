package database

import (
	"database/sql"
	"errors"
	"fmt"

	// Import the PostgreSQL driver.
	_ "github.com/lib/pq"
)

const (
	Driver                    = "postgres"
	MaxConnectionsDefault     = 5
	MaxIdleConnectionsDefault = 1
)

// Config represents the configuration for a database connection.
type Config struct {
	// DSN is the Data Source Name. It specifies the username, password, and database name
	// that are used to connect to the database.
	DSN string

	// MaxConnection is the maximum number of connections that can be opened to the database.
	MaxConnections int

	// MaxIdleConnections is the maximum number of idle connections that can be maintained.
	// Idle connections are connections that are open but not in use.
	MaxIdleConnections int
}

// NewConnection creates a new database connection using the provided configuration.
func NewConnection(cfg Config) (*sql.DB, error) {
	if cfg.DSN == "" {
		return nil, errors.New("data source name (DSN) is required")
	}

	db, err := sql.Open(Driver, cfg.DSN)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %w", err)
	}

	if cfg.MaxConnections <= 0 {
		db.SetMaxOpenConns(MaxConnectionsDefault)
	} else {
		db.SetMaxOpenConns(cfg.MaxConnections)
	}

	if cfg.MaxIdleConnections <= 0 {
		db.SetMaxIdleConns(MaxIdleConnectionsDefault)
	} else {
		db.SetMaxIdleConns(cfg.MaxIdleConnections)
	}

	return db, err
}
