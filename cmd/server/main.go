package main

import "os"

func main() {
	rootCmd := NewRootCMD()
	rootCmd.AddCommand(NewServeCMD().Command)

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
