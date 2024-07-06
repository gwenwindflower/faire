package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var inProgressStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#CA9EE6"))

var doneStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a6d189"))

var headerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#CA9EE6")).
	Padding(0, 0, 1, 0)

var selectedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#414559")).
	Background(lipgloss.Color("#C6D0F5"))

func (m model) ViewSelect() string {
	s := headerStyle.Render("What do you want to do today 🤗?") + "\n"
	var task string
	for i, todo := range m.todos {
		if m.hideCompleted && m.todos[i].Status == Done {
			continue
		}
		if m.cursor == i {
			task = selectedStyle.Render(todo.Task)
		} else {
			task = todo.Task
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
		if todo.DueDate.IsSet() {
			s += fmt.Sprintf("%s %s %s\n", checked, task, todo.DueDate.Format("2006-01-02"))
		} else {
			s += fmt.Sprintf("%s %s\n", checked, task)
		}
	}
	s += "\nPress 'a' to add a new todo, 'q' to quit."
	return s
}
