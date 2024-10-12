package main

import (
	"github.com/b-url/burl/cmd/server/config"
	"github.com/spf13/cobra"
)

// NewRootCMD returns a new root command.
func NewRootCMD() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "burlserver",
		Short: "burlserver is the server for the burl bookmark management tool",
	}

	cmd.PersistentFlags().String(
		config.FlagLogLevel,
		"INFO",
		"The log level to use. (DEBUG, INFO, WARN, ERROR)",
	)

	cmd.PersistentFlags().String(
		config.FlagLogType,
		"TEXT",
		"The type of log output (TEXT or JSON)",
	)

	return cmd
}
