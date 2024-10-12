package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
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

const (
	DefaultReadTimeout  = time.Second * 5
	DefaultWriteTimeout = time.Second * 5
)

type ServeCommand struct {
	Command *cobra.Command

	logger *slog.Logger
	config *config.Config
}

// NewServeCMD returns a new serve command.
func NewServeCMD() *ServeCommand {
	config := config.New()
	serveCommand := &ServeCommand{
		config: config,
		logger: config.NewLogger(),
	}

	serveCommand.Command = &cobra.Command{
		Use:   "serve",
		Short: "Serve the burl server",
		Long:  `This command starts the burl server.`,
		RunE:  serveCommand.Execute,
	}

	return serveCommand
}

// Execute creates the dependencies for the server and starts serving.
func (sc *ServeCommand) Execute(cmd *cobra.Command, _ []string) error {
	// Create a context that listens for the interrupt signal.
	ctx, stop := signal.NotifyContext(cmd.Context(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	dsn, err := sc.config.DBURL()
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
	bookmarker := bookmark.NewBookmarker(repo)

	return sc.Serve(ctx, apiimpl.NewServer(bookmarker))
}

// Serve starts the burl server and blocks until it's shut down.
func (sc *ServeCommand) Serve(ctx context.Context, server api.ServerInterface) error {
	p, err := sc.config.HTTPPort()
	if err != nil {
		return err
	}

	sc.logger.InfoContext(ctx, "Starting server", "port", p)

	r := http.NewServeMux()
	h := api.HandlerFromMux(server, r)
	s := &http.Server{
		Handler:      h,
		Addr:         fmt.Sprintf(":%d", p),
		ReadTimeout:  DefaultReadTimeout,
		WriteTimeout: DefaultWriteTimeout,
	}

	// Start the server in a goroutine.
	errChan := make(chan error, 1)
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			sc.logger.ErrorContext(ctx, "Server is closing", "error", err)
			errChan <- err
		}
	}()

	// Wait for context cancellation or server error.
	select {
	case <-ctx.Done():
		sc.logger.InfoContext(ctx, "Shutting down server")
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
