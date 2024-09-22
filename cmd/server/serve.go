package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	api "github.com/b-url/burl/api/v1"
	apiimpl "github.com/b-url/burl/cmd/server/api"
	"github.com/b-url/burl/cmd/server/config"
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
			_, err := Serve(cmd.Context(), config.New(), apiimpl.NewServer())
			return err
		},
	}
}

// Serve starts the burl server and returns a function to shut it down.
func Serve(_ context.Context, c *config.Config, server api.ServerInterface) (func(ctx context.Context) error, error) {
	p, err := c.HTTPPort()
	if err != nil {
		return nil, err
	}
	fmt.Printf("Starting server on port %d\n", p)
	r := http.NewServeMux()
	h := api.HandlerFromMux(server, r)
	s := &http.Server{
		Handler:     h,
		Addr:        fmt.Sprintf(":%d", p),
		ReadTimeout: readTimeout,
	}
	go func() {
		if err = s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()
	return s.Shutdown, nil
}
