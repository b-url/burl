package main

import (
	"context"
	"testing"

	"github.com/b-url/burl/cmd/server/config"
	"github.com/spf13/cobra"
)

func TestNewServeCMD(t *testing.T) {
	t.Run("should return a new serve command", func(t *testing.T) {
		if NewServeCMD() == nil {
			t.Error("NewServeCMD() = nil; want a new serve command")
		}
	})

	t.Run("RunE should call Serve", func(t *testing.T) {
		t.Setenv("BURLSERVER_DB_URL", "postgres://localhost:5432/burl")
		t.Setenv("BURLSERVER_HTTP_PORT", "8080")
		err := NewServeCMD().RunE(&cobra.Command{}, []string{})
		if err != nil {
			t.Errorf("NewServeCMD().RunE() error = %v; want nil", err)
		}
	})
}

func TestServe(t *testing.T) {
	t.Run("should start the server", func(t *testing.T) {
		c := config.New()
		t.Setenv("BURLSERVER_DB_URL", "postgres://localhost:5432/burl")
		t.Setenv("BURLSERVER_HTTP_PORT", "8080")
		err := Serve(context.TODO(), c)
		if err != nil {
			t.Errorf("Serve() error = %v; want nil", err)
		}
	})

	t.Run("should return an error if the DB URL is not set", func(t *testing.T) {
		c := config.New()
		t.Setenv("BURLSERVER_HTTP_PORT", "8080")
		err := Serve(context.TODO(), c)
		if err == nil {
			t.Error("Serve() error = nil; want an error")
		}
	})

	t.Run("should return an error if the HTTP port is not set", func(t *testing.T) {
		c := config.New()
		t.Setenv("BURLSERVER_DB_URL", "postgres://localhost:5432/burl")
		err := Serve(context.TODO(), c)
		if err == nil {
			t.Error("Serve() error = nil; want an error")
		}
	})
}
