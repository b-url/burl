package main

import (
	"github.com/spf13/cobra"
)

// NewRootCMD returns a new root command.
func NewRootCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "burlserver",
		Short: "b(ookmark)url is a developer first bookmark management tool written in Go",
	}
}
