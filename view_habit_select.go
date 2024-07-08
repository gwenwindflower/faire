package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	habitDateStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#A9B1D6")).
			Width(appWidth).
			Align(lipgloss.Center)

	habitCheckedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#a6d189"))

	habitMissedStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#e78284"))

	selectedHabitNameStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#414559")).
				Background(lipgloss.Color("#C6D0F5"))

	activeHabitDayStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#414559")).
				Background(lipgloss.Color("#C6D0F5"))
)

func (m model) ViewHabitSelect() string {
	s := ""
	header := headerStyle.Render("Habit Tracker") + "\n"
	s += header
	date := habitDateStyle.Render(m.aWeekBeforeActiveHabitDay.Format("Mon Jan 2"), " - ", m.activeHabitDay.Format("Mon Jan 2"))
	s += date + "\n"

	t := table.New().
		Width(appWidth).
		Headers(
			"Habit",
			"Days",
		)
	for i, name := range m.habitList {
		if m.habitCursor == i {
			t.Row(selectedHabitNameStyle.Render(name), renderHabitWeek((*m.habits)[name], m, true))
		} else {
			t.Row(name, renderHabitWeek((*m.habits)[name], m, false))
		}
	}
	t.StyleFunc(func(row, col int) lipgloss.Style {
		if col == 0 {
			return lipgloss.NewStyle().
				Width(54).
				Align(lipgloss.Left)
		} else {
			return lipgloss.NewStyle().
				Width(12).
				Align(lipgloss.Center)
		}
	})
	s += t.Render() + "\n"
	footer := footerStyle.Render("Press '?' for shortcuts.")
	s += footer
	return s
}

func renderHabitWeek(habitDays []Habit, m model, selected bool) string {
	s := ""
	for _, day := range habitDays {
		if day.Date.After(m.aWeekBeforeActiveHabitDay) && (day.Date.Before(m.activeHabitDay) || day.Date.Equal(m.activeHabitDay)) {
			if day.Completed {
				if selected && day.Date.Equal(m.activeHabitDay) {
					s += activeHabitDayStyle.Render("✔")
				} else {
					s += habitCheckedStyle.Render("✔")
				}
			} else {
				if selected && day.Date.Equal(m.activeHabitDay) {
					s += activeHabitDayStyle.Render("-")
				} else {
					s += habitMissedStyle.Render("-")
				}
			}
		}
	}
	return s
}
