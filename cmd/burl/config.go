package main

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/b-url/burl/cmd/burl/config"
	"github.com/spf13/cobra"
)

type ConfigCommand struct {
	command *cobra.Command
}

func NewConfigCommand() *ConfigCommand {
	cmd := &ConfigCommand{}
	cmd.command = &cobra.Command{
		Use:   "config",
		Short: "Manage the configuration of the burl command",
	}

	cmd.command.AddCommand(editConfigCommand)
	return cmd
}

var editConfigCommand = &cobra.Command{
	Use:   "edit",
	Short: "Edit the configuration file",
	RunE: func(_ *cobra.Command, _ []string) error {
		config, err := config.New()
		if err != nil {
			return err
		}

		filepath, err := config.Write()
		if err != nil {
			return err
		}

		editorName := os.Getenv("EDITOR")
		if runtime.GOOS == "windows" {
			editorName = "notepad"
		}
		if editorName == "" {
			editorName = "vi"
		}
		cmd := exec.Command(editorName, filepath)
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout

		if err = cmd.Run(); err != nil {
			return err
		}

		return nil
	},
}
