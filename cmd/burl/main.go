package main

import "os"

func main() {
	rootCmd := NewRootCommand()

	if err := rootCmd.command.Execute(); err != nil {
		os.Exit(1)
	}
}
