package main

import (
	// sqllite3 driver
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	rootCmd := NewRootCommand()
	addCmd := NewAddCommand()

	rootCmd.command.AddCommand(addCmd.command)

	if err := rootCmd.command.Execute(); err != nil {
		os.Exit(1)
	}
}
