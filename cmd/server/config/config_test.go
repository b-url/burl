package config_test

import (
	"testing"

	"github.com/b-url/burl/cmd/server/config"
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
		t.Setenv("BURL_DB_URL", "postgres://localhost:5432/burl")
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
		t.Setenv("BURL_HTTP_PORT", "8080")
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
