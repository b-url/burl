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

func (m RootTui) Update(_ tea.Msg) (tea.Model, tea.Cmd) {
	return m, nil
}

func (m RootTui) View() string {
	return "Hello, World!"
}
