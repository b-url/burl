package main

import "os"

func main() {
	rootCmd := NewRootCMD()
	rootCmd.AddCommand(NewServeCMD())

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
