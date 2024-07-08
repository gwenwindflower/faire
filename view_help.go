package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var helpHeaderStyle = headerStyle.
	Width(appWidth).
	Border(lipgloss.RoundedBorder()).
	Align(lipgloss.Center)

var subheaderStyle = lipgloss.NewStyle().
	Width(appWidth).
	Align(lipgloss.Center).
	Foreground(lipgloss.Color("#a6d189"))

func (m model) ViewHelp() string {
	header := helpHeaderStyle.Render("Shortcuts") + "\n"
	subHeader := subheaderStyle.Render(func() string {
		switch m.activeHelpScreenId {
		case SelectScreenId:
			return "Todo List"
		case AddScreenId:
			return "Add Todo"
		case HabitScreenId:
			return "Habit Tracker"
		case AddHabitScreenId:
			return "Add Habit"
		}
		return "Oops! Something bad happened"
	}()) + "\n"

	t := table.New().
		Headers("Key(s)", "Action").
		Width(appWidth).
		StyleFunc(func(row, col int) lipgloss.Style {
			if col == 0 {
				style := lipgloss.NewStyle().Width(16)
				if row != 0 {
					style = style.Foreground(lipgloss.Color("#f4b8e4"))
				}
				return style
			}
			return lipgloss.NewStyle().Width(50)
		})

	switch m.activeHelpScreenId {

	case SelectScreenId:
		t.Rows([][]string{
			{"space", "Cycle todo status"},
			{"1/2/3", "Set todo status"},
			{"a", "Add todo"},
			{"d", "Delete todo"},
			{"D", "Delete all completed"},
			{"h", "Toggle show completed"},
			{"u", "Undo delete"},
			{"j/k", "Move cursor"},
			{"tab/shift+tab", "Move cursor"},
			{"down/up", "Move cursor"},
			{"ctrl+n/ctrl+p", "Move cursor"},
			{"q/ctrl+c", "Quit"},
		}...)
	case AddScreenId:
		t.Rows([][]string{
			{"esc", "Return to list"},
			{"enter", "Save todo"},
			{"tab/shift+tab", "Move input focus"},
			{"up/down", "Move input focus"},
			{"ctrl+c", "Quit"},
		}...)
	case HabitScreenId:
		t.Rows([][]string{
			{"space", "Cycle habit status"},
			{"a", "Add habit"},
			{"h/l", "Move active day"},
			{"H/L", "Move active day by week"},
			{"j/k", "Move cursor"},
			{"down/up", "Move cursor"},
			{"ctrl+n/ctrl+p", "Move cursor"},
			{"tab/shift+tab", "Move cursor"},
			{"q/ctrl+c", "Quit"},
		}...)
	case AddHabitScreenId:
		t.Rows([][]string{
			{"enter", "Save habit"},
			{"esc", "Return to habit tracker"},
			{"ctrl+c", "Quit without saving"},
		}...)
	}
	footer := footerStyle.Render("\nUse h/l to cycle through help screens.")
	s := header
	s += subHeader
	s += t.Render()
	s += footer
	return s
}
