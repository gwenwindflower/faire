package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

const appWidth = 66

var (
	cuteBorder = lipgloss.Border{
		Bottom: "🌿🌸🫧✨🦋",
	}

	inProgressStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CA9EE6"))

	doneStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#a6d189"))

	headerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#CA9EE6")).
			Width(appWidth).
			Align(lipgloss.Center).
			Bold(true).
			BorderStyle(cuteBorder).BorderBottom(true)

	footerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#838ba7")).
			Width(appWidth).
			Align(lipgloss.Center)

	todayStyle = footerStyle

	selectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#414559")).
			Background(lipgloss.Color("#C6D0F5"))

	dueDateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#f4b8e4")).
			Align(lipgloss.Right)
)

func (m model) ViewSelect() string {
	t := table.New().
		Headers("Status", "Task", "Due Date").
		Width(appWidth).
		StyleFunc(func(row, col int) lipgloss.Style {
			switch col {
			case 0:
				return lipgloss.NewStyle().Width(8).Align(lipgloss.Center)
			case 1:
				if row == 0 {
					return lipgloss.NewStyle().Width(48).Align(lipgloss.Center)
				} else {
					return lipgloss.NewStyle().Width(48).Align(lipgloss.Left)
				}
			case 2:
				return lipgloss.NewStyle().Width(12).Align(lipgloss.Center)
			default:
				return lipgloss.NewStyle()
			}
		})

	s := headerStyle.Render("What do you want to do today 🤗?") + "\n"
	s += todayStyle.Render(m.today.Format("Jan 2, 2006")) + "\n"
	var task string
	for i, todo := range *m.todos {
		if m.hideCompleted && (*m.todos)[i].Status == Done {
			continue
		}
		if m.todoCursor == i {
			task = " " + selectedStyle.Render(todo.Task)
		} else {
			task = " " + todo.Task
		}
		checked := " "
		switch (*m.todos)[i].Status {
		case InProgress:
			checked = inProgressStyle.Render("󰦕")
		case Done:
			checked = doneStyle.Render("󰗡")
		default:
			checked = "󰄰"
		}
		if todo.DueDate.IsSet() {
			dueDate := dueDateStyle.Render(todo.DueDate.Format("2006-01-02"))
			t.Row(checked, task, dueDate)
		} else {
			t.Row(checked, task, "")
		}
	}
	s += t.Render() + "\n"
	if m.hideCompleted {
		s += subheaderStyle.
			Width(appWidth).
			Align(lipgloss.Center).
			Render("Completed tasks hidden.\n")
	}
	footer := footerStyle.Render("\nPress ? for shortcuts.")
	s += footer
	return s
}
