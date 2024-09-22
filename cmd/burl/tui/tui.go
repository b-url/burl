// Package tui provides the terminal user interface for the burl command.
// It uses the bubbletea library to render the UI.
package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Force Model to be a tea.Model.
var _ tea.Model = RootTui{}

type RootTui struct {
	keys keyMap

	height int
	width  int

	// Components
	help      help.Model
	table     *Table
	statusBar *StatusBar
}

func New() *RootTui {
	return &RootTui{
		keys:      keys,
		help:      help.New(),
		table:     NewTable(),
		statusBar: NewStatusBar("currentCollection"),
	}
}

func (m RootTui) Init() tea.Cmd {
	return nil
}

func (m RootTui) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	cmds := []tea.Cmd{}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit

		case key.Matches(msg, m.keys.Open):
			// TODO: Open URL
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

		m.table.SetSize(msg.Width)
		m.statusBar.SetSize(msg.Width)
	}

	cmds = append(cmds, m.table.Update(msg))
	return m, tea.Batch(cmds...)
}

func (m RootTui) View() string {
	help := m.help.View(keys)
	statusBar := m.statusBar.View()

	helpBox := lipgloss.NewStyle().
		Width(m.width).
		Height(lipgloss.Height(help)).
		MarginBottom(1).
		Render(help)

	bodyBox := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Width(m.width - 2).
		Height(m.height - lipgloss.Height(statusBar) - lipgloss.Height(helpBox) - 2).
		Render(m.table.View())

	return lipgloss.JoinVertical(
		lipgloss.Top,
		bodyBox,
		helpBox,
		statusBar,
	)
}
