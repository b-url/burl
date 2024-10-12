package config_test

import (
	"log/slog"
	"testing"

	"github.com/b-url/burl/cmd/server/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	t.Run("should return a new Config", func(t *testing.T) {
		if config.New() == nil {
			t.Error("New() = nil; want a new Config")
		}
	})
}

func TestConfig_DBURL(t *testing.T) {
	t.Run("should return the DB URL", func(t *testing.T) {
		c := config.New()
		t.Setenv("BURLSERVER_DB_URL", "postgres://localhost:5432/burl")
		got, err := c.DBURL()
		if err != nil {
			t.Errorf("DBURL() error = %v; want nil", err)
		}
		if got != "postgres://localhost:5432/burl" {
			t.Errorf("DBURL() = %s; want postgres://localhost:5432/burl", got)
		}
	})

	t.Run("should return an error if the DB URL is not set", func(t *testing.T) {
		c := config.New()
		_, err := c.DBURL()
		if err == nil {
			t.Error("DBURL() error = nil; want an error")
		}
	})
}

func TestConfig_HTTPPort(t *testing.T) {
	t.Run("should return the HTTP port", func(t *testing.T) {
		c := config.New()
		t.Setenv("BURLSERVER_HTTP_PORT", "8080")
		got, err := c.HTTPPort()
		if err != nil {
			t.Errorf("HTTPPort() error = %v; want nil", err)
		}
		if got != 8080 {
			t.Errorf("HTTPPort() = %d; want 8080", got)
		}
	})

	t.Run("should return an error if the HTTP port is not set", func(t *testing.T) {
		c := config.New()
		_, err := c.HTTPPort()
		if err == nil {
			t.Error("HTTPPort() error = nil; want an error")
		}
	})
}

func TestConfig_LogLevel(t *testing.T) {
	tests := []struct {
		name      string
		envValue  string
		wantLevel slog.Level
		wantErr   bool
	}{
		{"valid log level info", "info", slog.LevelInfo, false},
		{"valid log level debug", "debug", slog.LevelDebug, false},
		{"invalid log level", "invalid", slog.LevelInfo, true},
		{"empty log level", "", slog.LevelInfo, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Setenv("BURLSERVER_LOG_LEVEL", tt.envValue)

			cfg := config.New()
			gotLevel, err := cfg.LogLevel()

			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			assert.Equal(t, tt.wantLevel, gotLevel)
		})
	}
}
