package main

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

var habitCheckedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#a6d189"))

var habitMissedStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#e78284"))

var selectedHabitNameStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#414559")).
	Background(lipgloss.Color("#C6D0F5"))

func (m model) ViewHabitSelect() string {
	s := ""
	header := headerStyle.Render("Habits")
	s += header + "\n"

	// Calculate the cutoff time once
	cutoffTime := time.Now().Add(-7 * 24 * time.Hour)

	for i, name := range m.habitList {
		if m.habitCursor == i {
			s += selectedHabitNameStyle.Render(name) + " "
		} else {
			s += name + " "
		}
		for _, day := range m.habits[name] {
			if day.Date.After(cutoffTime) || day.Date.Equal(cutoffTime) {
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
