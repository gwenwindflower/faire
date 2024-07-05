package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var cursorStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#c6d0f5"))

var inProgressStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#CA9EE6"))

var doneStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a6d189"))

func (m model) ViewSelect() string {
	s := "What do you want to accomplish?\n\n"
	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = "➔"
		}
		checked := " "
		switch m.todos[i].Status {
		case InProgress:
			checked = inProgressStyle.Render("◌")
		case Done:
			checked = doneStyle.Render("✔")
		default:
			checked = "⬚"
		}
		s += fmt.Sprintf("%s %s %s\n", cursorStyle.Render(cursor), checked, choice)
	}
	s += "\nPress 'a' to add a new todo, 'q' to quit."
	return s
}
