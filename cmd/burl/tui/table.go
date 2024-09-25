package tui

import (
	"os/exec"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var baseStyle = lipgloss.NewStyle().BorderForeground(lipgloss.Color("240"))

type Table struct {
	width int
	table table.Model
}

func NewTable() *Table {
	columns := []table.Column{
		{Title: "Name", Width: 20},
		{Title: "URL", Width: 40},
	}

	t := table.New(
		table.WithColumns(columns),
		table.WithRows([]table.Row{
			// TODO(mveen): Remove mocks and inject properly.
			{"Google", "https://google.com"},
			{"DuckDuckGo", "https://duckduckgo.com"},
			{"HackerNews", "https://news.ycombinator.com"},
			{"GitHub", "https://github.com"},
			{"YouTube", "https://youtube.com"},
		}),
		table.WithFocused(true),
		table.WithHeight(7),
	)

	s := table.DefaultStyles()
	s.Header = s.Header.
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("240")).
		BorderBottom(true).
		Bold(false)

	s.Selected = s.Selected.
		Foreground(lipgloss.Color("229")).
		Background(lipgloss.Color("57")).
		Bold(false)
	t.SetStyles(s)

	return &Table{
		table: t,
	}
}

func (t *Table) SetSize(width int) {
	t.width = width
	t.table.SetWidth(width)
}

func (t *Table) Update(msg tea.Msg) tea.Cmd {
	var cmd tea.Cmd
	t.table, cmd = t.table.Update(msg)
	return cmd
}

func (t *Table) View() string {
	return baseStyle.Render(t.table.View())
}

func (t *Table) OpenSelected() tea.Cmd {
	selected := t.table.SelectedRow()

	c := exec.Command("open", selected[1]) //nolint:gosec // This is a CLI tool
	return tea.ExecProcess(c, func(err error) tea.Msg {
		if err != nil {
			return err.Error()
		}
		return nil
	})
}
