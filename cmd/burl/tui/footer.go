package tui

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/lipgloss"
)

var (
	HelpStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	ColorOrange = lipgloss.Color("172")
)

type Footer struct {
	keys              keyMap
	help              help.Model
	currentCollection string
}

func NewFooter(keys keyMap, initCollection string) *Footer {
	return &Footer{currentCollection: initCollection, keys: keys, help: help.New()}
}

// Update updates the footer.
func (f *Footer) Update(collection string) {
	f.currentCollection = collection
}

// View returns the footer view.
func (f *Footer) View() string {
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
