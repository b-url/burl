package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	api "github.com/b-url/burl/api/v1"
	apiimpl "github.com/b-url/burl/cmd/server/api"
	"github.com/b-url/burl/cmd/server/bookmark"
	"github.com/b-url/burl/cmd/server/config"
	"github.com/b-url/burl/cmd/server/database"
	"github.com/spf13/cobra"
)

const readTimeout = time.Second * 5

// NewServeCMD returns a new serve command.
func NewServeCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Serve the burl server",
		Long:  `This command starts the burl server.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			// Create a context that listens for the interrupt signal.
			ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
			defer stop()

			cfg := config.New()
			dsn, err := cfg.DBURL()
			if err != nil {
				return err
			}

			db, err := database.NewConnection(database.Config{
				DSN: dsn,
			})
			if err != nil {
				return err
			}

			repo := bookmark.NewRepository(db)
			b := bookmark.NewBookmarker(repo)

			return Serve(ctx, config.New(), apiimpl.NewServer(b))
		},
	}
}

// Serve starts the burl server and blocks until it's shut down.
func Serve(ctx context.Context, c *config.Config, server api.ServerInterface) error {
	p, err := c.HTTPPort()
	if err != nil {
		return err
	}
	fmt.Printf("Starting server on port %d\n", p)
	r := http.NewServeMux()
	h := api.HandlerFromMux(server, r)
	s := &http.Server{
		Handler:     h,
		Addr:        fmt.Sprintf(":%d", p),
		ReadTimeout: readTimeout,
	}

	// Start the server in a goroutine.
	errChan := make(chan error, 1)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errChan <- err
		}
	}()

	// Wait for context cancellation or server error.
	select {
	case <-ctx.Done():
		// Shutdown the server gracefully.
		ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err = s.Shutdown(ctxShutdown); err != nil {
			return err
		}
	case err = <-errChan:
		return err
	}

	return nil
}
