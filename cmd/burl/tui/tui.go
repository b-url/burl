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
	table  *Table
	footer *StatusBar
}

func New() *RootTui {
	return &RootTui{
		keys:   keys,
		table:  NewTable(),
		footer: NewStatusBar(keys, "currentCollection"),
	}
}

func (m RootTui) Init() tea.Cmd {
	return nil
}

func (m RootTui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	//nolint:gocritic // Will be expanded in the future
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		}
	}

	cmds = append(cmds, m.table.Update(msg))
	return m, tea.Batch(cmds...)
}

func (m RootTui) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Left,
		m.table.View(),
		m.footer.View(),
	)
}
