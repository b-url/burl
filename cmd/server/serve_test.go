package main

import (
	"context"
	"testing"

	apiimpl "github.com/b-url/burl/cmd/server/api"
	"github.com/b-url/burl/cmd/server/config"
)

func TestNewServeCMD(t *testing.T) {
	t.Run("should return a new serve command", func(t *testing.T) {
		if NewServeCMD() == nil {
			t.Error("NewServeCMD() = nil; want a new serve command")
		}
	})
}

func TestServe(t *testing.T) {
	t.Run("should start the server", func(t *testing.T) {
		c := config.New()
		t.Setenv("BURLSERVER_DB_URL", "postgres://localhost:5432/burl")
		t.Setenv("BURLSERVER_HTTP_PORT", "7777")
		shutdown, err := Serve(context.TODO(), c, apiimpl.NewServer())
		if err != nil {
			t.Errorf("Serve() error = %v; want nil", err)
		}
		if shutdown == nil {
			t.Error("Serve() shutdown = nil; want a function")
		}
		if err = shutdown(context.Background()); err != nil {
			t.Errorf("shutdown() error = %v; want nil", err)
		}
	})

	t.Run("should return an error if the HTTP port is not set", func(t *testing.T) {
		c := config.New()
		t.Setenv("BURLSERVER_DB_URL", "postgres://localhost:5432/burl")
		_, err := Serve(context.TODO(), c, apiimpl.NewServer())
		if err == nil {
			t.Error("Serve() error = nil; want an error")
		}
	})
}
