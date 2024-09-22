// Package tui provides the terminal user interface for the burl command.
// It uses the bubbletea library to render the UI.
package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Force Model to be a tea.Model.
var _ tea.Model = RootTui{}

type RootTui struct{}

func New() *RootTui {
	return &RootTui{}
}

func (m RootTui) Init() tea.Cmd {
	return nil
}

func (m RootTui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic // Will be expanded in the future
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m RootTui) View() string {
	return "Hello, World!"
}
