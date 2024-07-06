package main

import tea "github.com/charmbracelet/bubbletea"

func (m model) UpdateHelp(msg tea.Msg) (model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "esc", "?":
			m.activeView = m.previousViewFromHelp
		}
	}
	return m, nil
}
