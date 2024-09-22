package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

var (
	HelpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	ColorOrange = lipgloss.Color("172")
)

type StatusBar struct {
	keys              keyMap
	help              help.Model
	currentCollection string
}

func NewStatusBar(keys keyMap, initCollection string) *StatusBar {
	return &StatusBar{currentCollection: initCollection, keys: keys, help: help.New()}
}

// Update updates the footer.
func (f *StatusBar) Update(collection string) {
	f.currentCollection = collection
}

// View returns the footer view.
func (f *StatusBar) View() string {
	views := []string{}

	views = append(views, f.help.View(keys))
	currentCollection := lipgloss.NewStyle().
		Margin(0, 1).
		Padding(0, 1).
		Bold(true).
		Background(ColorOrange).
		Render(f.currentCollection)

	views = append(views, currentCollection)

	return lipgloss.JoinHorizontal(lipgloss.Left, views...)
}
