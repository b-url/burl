package main

import (
	"context"
	"testing"
	"time"

	apiimpl "github.com/b-url/burl/cmd/server/api"
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
		serveCommand := NewServeCMD()

		t.Setenv("BURLSERVER_DB_URL", "postgres://localhost:5432/burl")
		t.Setenv("BURLSERVER_HTTP_PORT", "7777")

		// Create a context with cancel
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		// Run the server in a separate goroutine to simulate actual use
		errChan := make(chan error, 1)
		go func() {
			errChan <- serveCommand.Serve(ctx, apiimpl.NewServer(nil))
		}()

		// Give the server some time to start
		time.Sleep(10 * time.Millisecond)

		// Check if Serve returned an error
		select {
		case err := <-errChan:
			if err != nil {
				t.Errorf("Serve() error = %v; want nil", err)
			}
		default:
			// No error yet, so assume it's running fine
		}

		// Now, simulate server shutdown
		cancel()

		// Wait for Serve to return
		select {
		case err := <-errChan:
			if err != nil {
				t.Errorf("Serve() error during shutdown = %v; want nil", err)
			}
		case <-time.After(2 * time.Second):
			t.Error("Serve() did not return within 2 seconds after shutdown; likely hung")
		}
	})

	t.Run("should return an error if the HTTP port is not set", func(t *testing.T) {
		serveCommand := NewServeCMD()

		t.Setenv("BURLSERVER_DB_URL", "postgres://localhost:5432/burl")

		// Use a context with a short timeout
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		err := serveCommand.Serve(ctx, apiimpl.NewServer(nil))
		if err == nil {
			t.Error("Serve() error = nil; want an error")
		}
	})
}
