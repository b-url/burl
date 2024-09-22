package main

import (
	"context"
	"fmt"

	"github.com/b-url/burl/cmd/server/config"
	"github.com/spf13/cobra"
)

// NewServeCMD returns a new serve command.
func NewServeCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "serve",
		Short: "Serve the burl server",
		Long:  `This command starts the burl server.`,
		RunE: func(cmd *cobra.Command, _ []string) error {
			return Serve(cmd.Context(), config.New())
		},
	}
}

// Serve starts the burl server.
func Serve(_ context.Context, c *config.Config) error {
	p, err := c.HTTPPort()
	if err != nil {
		return err
	}
	fmt.Printf("starting server on port %d\n", p)
	d, err := c.DBURL()
	if err != nil {
		return err
	}
	fmt.Printf("connecting to database at %s\n", d)
	return nil
}
