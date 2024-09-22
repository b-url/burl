package main

import (
	"github.com/b-url/burl/cmd/burl/tui"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

// NewRootCMD returns a new root command.
func NewRootCMD() *cobra.Command {
	return &cobra.Command{
		Use:   "burl",
		Short: "b(ookmark)url is a developer first bookmark management tool written in Go",
		RunE: func(_ *cobra.Command, _ []string) error {
			p := tea.NewProgram(tui.New(), tea.WithAltScreen())
			_, err := p.Run()
			return err
		},
	}
}
