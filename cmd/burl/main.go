package main

import "os"

func main() {
	rootCmd := NewRootCMD()

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
