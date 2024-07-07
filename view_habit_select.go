package main

import (
	"github.com/charmbracelet/lipgloss"
)

var habitDateStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#A9B1D6"))

var habitCheckedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a6d189"))

var habitMissedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#e78284"))

var selectedHabitNameStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#414559")).
	Background(lipgloss.Color("#C6D0F5"))

func (m model) ViewHabitSelect() string {
	s := ""
	header := headerStyle.Render("Habit Tracker") + "\n"
	s += header
	date := habitDateStyle.Render(m.aWeekBeforeActiveHabitDay.Format("Mon Jan 2"), " - ", m.activeHabitDay.Format("Mon Jan 2"))
	s += date + "\n"
	for i, name := range m.habitList {
		if m.habitCursor == i {
			s += selectedHabitNameStyle.Render(name) + " "
		} else {
			s += name + " "
		}
		for _, day := range m.habits[name] {
			if (day.Date.After(m.aWeekBeforeActiveHabitDay) && day.Date.Before(m.activeHabitDay)) || day.Date.Equal(m.activeHabitDay) {
				if day.Completed {
					s += habitCheckedStyle.Render("✔")
				} else {
					s += habitMissedStyle.Render("-")
				}
			}
		}
		s += "\n"
	}
	footer := footerStyle.Render("Press '?' for shortcuts.")
	s += footer
	return s
}
