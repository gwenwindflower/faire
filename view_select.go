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
	Bold(true)

var footerStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#838ba7"))

var selectedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#414559")).
	Background(lipgloss.Color("#C6D0F5"))

var dueDateStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#f4b8e4"))

func (m model) ViewSelect() string {
	s := headerStyle.Render("What do you want to do today 🤗?") + "\n"
	var task string
	for i, todo := range *m.todos {
		if m.hideCompleted && (*m.todos)[i].Status == Done {
			continue
		}
		if m.todoCursor == i {
			task = selectedStyle.Render(todo.Task)
		} else {
			task = todo.Task
		}
		checked := " "
		switch (*m.todos)[i].Status {
		case InProgress:
			checked = inProgressStyle.Render("◌")
		case Done:
			checked = doneStyle.Render("✔")
		default:
			checked = "⬚"
		}
		if todo.DueDate.IsSet() {
			dueDate := dueDateStyle.Render(todo.DueDate.Format("2006-01-02"))
			s += fmt.Sprintf("%s %s %s\n", checked, task, dueDate)
		} else {
			s += fmt.Sprintf("%s %s\n", checked, task)
		}
	}
	footer := footerStyle.Render("Press '?' for shortcuts.")
	s += footer
	return s
}
