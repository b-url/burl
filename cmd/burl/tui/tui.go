// Package tui provides the terminal user interface for the burl command.
// It uses the bubbletea library to render the UI.
package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Force Model to be a tea.Model.
var _ tea.Model = RootTui{}

type RootTui struct {
	keys   keyMap
	footer *Footer
}

func New() *RootTui {
	return &RootTui{
		keys:   keys,
		footer: NewFooter(keys, "currentCollection"),
	}
}

func (m RootTui) Init() tea.Cmd {
	return nil
}

func (m RootTui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	//nolint:gocritic // Will be expanded in the future
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m RootTui) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.footer.View(),
	)
}
