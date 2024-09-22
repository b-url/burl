package main

import "os"

func main() {
	rootCmd := NewRootCMD()
	rootCmd.AddCommand(NewServeCMD())
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
