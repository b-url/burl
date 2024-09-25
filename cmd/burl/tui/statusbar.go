package tui

import (
	"github.com/charmbracelet/lipgloss"
)

const StatusBarHeight = 1

var (
	HelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))

	collectionBackground = lipgloss.Color("164")
	versionBackground    = lipgloss.Color("45")
)

type StatusBar struct {
	currentCollection string
	version           string

	height int
	width  int
}

func NewStatusBar(initCollection string) *StatusBar {
	return &StatusBar{currentCollection: initCollection, version: "v0.0.1-alpha"} // TODO: get version from somewhere
}

func (f *StatusBar) Update(collection string) {
	f.currentCollection = collection
}

func (f *StatusBar) SetSize(width int) {
	f.width = width
}

func (f *StatusBar) View() string {
	width := lipgloss.Width

	currentCollection := lipgloss.NewStyle().
		Padding(0, 1).
		Bold(true).
		Background(collectionBackground).
		Render(f.currentCollection)

	version := lipgloss.NewStyle().
		Padding(0, 1).
		Background(versionBackground).
		Render(f.version)

	space := lipgloss.NewStyle().
		Background(lipgloss.Color("241")).
		Padding(0, 1).
		Height(f.height).
		Width(f.width - width(currentCollection) - width(version)).
		Render("")

	return lipgloss.JoinHorizontal(lipgloss.Top, currentCollection, space, version)
}
