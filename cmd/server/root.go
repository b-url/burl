package main

import (
	"github.com/spf13/cobra"
)

// NewRootCMD returns a new root command.
func NewRootCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "burlserver",
		Short: "burlserver is the server for the burl bookmark management tool",
	}
}
