package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func (m model) UpdateHabitSelect(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down", "ctrl+n", "tab":
			if m.habitCursor < len(m.habitList)-1 {
				m.habitCursor++
			} else {
				m.habitCursor = 0
			}
		case "k", "up", "shift+tab", "ctrl+p":
			if m.habitCursor > 0 {
				m.habitCursor--
			} else {
				m.habitCursor = len(m.habitList) - 1
			}
		case "h", "left":
			m = m.MoveActiveHabitDay(-1)
		case "l", "right":
			m = m.MoveActiveHabitDay(1)
		case "H":
			m = m.MoveActiveHabitDay(-7)
		case "L":
			m = m.MoveActiveHabitDay(7)
		case " ":
			m = ToggleHabitToday(m)
		case "t":
			m.activeView = SelectViewId
		case "s":
			m.activeHabitDay = m.today
		case "a":
			m.addHabitInput.Focus()
			m.activeView = HabitAddViewId
		case "q", "ctrl+c":
			err := WriteAppData(m.dataFilePath, m.appData)
			if err != nil {
				log.Fatalf("Could not write app data: %v", err)
			}
			return m, tea.Quit
		case "?":
			m.previousViewFromHelp = HabitSelectViewId
			m.activeView = HelpViewId
		}
	}
	return m, nil
}
