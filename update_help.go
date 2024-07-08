package main

import tea "github.com/charmbracelet/bubbletea"

type HelpScreenId int

const (
	SelectScreenId HelpScreenId = iota
	AddScreenId
	HabitScreenId
	AddHabitScreenId
	MaxHelpScreenId
)

func (m model) UpdateHelp(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "h":
			// We use modulo to wrap around the help screens
			// The MaxHelpScreenId is just used to designate the cutoff/modulo point
			// by representing the number of help screens
			m.activeHelpScreenId = (m.activeHelpScreenId - 1 + MaxHelpScreenId) % MaxHelpScreenId
		case "l": // Increment
			m.activeHelpScreenId = (m.activeHelpScreenId + 1) % MaxHelpScreenId
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc", "?":
			m.activeView = m.previousViewFromHelp
		}
	}
	return m, nil
}
